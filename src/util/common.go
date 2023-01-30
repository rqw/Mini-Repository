package util

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoadConfig() *Config {
	return config
}
func GetFileSystem() http.FileSystem {
	return httpFs
}
func GetFileServer() http.Handler {
	return fileServer
}
func CreateParentIfNotExist(file string) error {
	dirPath := path.Dir(file)

	if stat, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			return err
		}
	} else if !stat.IsDir() {
		return fmt.Errorf("%s is not a dir", dirPath)
	}
	return nil
}

func CreateFileIfNotExist(file string) error {
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		if err = CreateParentIfNotExist(file); err != nil {
			return err
		}
		if _, err := os.Create(file); err != nil {
			return err
		}
	}
	return nil
}

func CheckFileExist(file string) (bool, error) {
	if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func GetHash(file []byte, hash string) []byte {
	switch hash {
	case "md5":
		return []byte(fmt.Sprintf("%x", md5.Sum(file)))
	case "sha1":
		return []byte(fmt.Sprintf("%x", sha1.Sum(file)))
	case "sha256":
		return []byte(fmt.Sprintf("%x", sha256.Sum256(file)))
	case "sha512":
		return []byte(fmt.Sprintf("%x", sha512.Sum512(file)))
	default:
		return nil
	}
}
func touchFile(file string, hash string, bytes []byte) error {
	hashFile := fmt.Sprintf("%s.%s", file, hash)
	if exist, err := CheckFileExist(hashFile); err != nil {
		return err
	} else if !exist {
		if err = os.WriteFile(hashFile, GetHash(bytes, hash), 0755); err != nil {
			return err
		}
	}
	return nil
}
func GenerateHash(file string) error {
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		dir, err := os.ReadDir(file)
		if err != nil {
			return err
		}
		for _, info := range dir {
			if err = GenerateHash(info.Name()); err != nil {
				return err
			}
		}
	}
	ext := path.Ext(file)
	if ext != ".xml" && ext != ".jar" && ext != ".pom" {
		return nil
	}
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	if err = touchFile(file, "md5", bytes); err != nil {
		return err
	}
	if err = touchFile(file, "sha1", bytes); err != nil {
		return err
	}
	return nil
}
func Md5(data string) string {
	hash := GetHash([]byte(data), "md5")
	return string(hash)
}
func JsonFileToAny(file string, target any) error {
	data, err := os.ReadFile(file)
	if err == nil {
		err = json.Unmarshal(data, &target)
	}
	return err
}
func AnyToJsonFile(target any, file string) error {
	content, err := json.Marshal(target)
	if err != nil {
		return err
	}
	if isExists, _ := CheckFileExist(config.DataDir); !isExists {
		os.MkdirAll(config.DataDir, os.ModePerm)
	}
	return os.WriteFile(file, content, os.ModePerm)
}
func GetParamId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, FAIL(err.Error(), nil))
	}
	return id, err
}
func GetParamJson[T any](c *gin.Context) (T, error) {
	var t T
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusOK, FAIL(err.Error(), nil))
	}
	return t, err
}

// ReleaseToken 签发token，id用户id字段，act用户act字段，expiresTime有效期（单位秒）
func ReleaseToken(id int, act string, expiresTime int64) string {
	content := fmt.Sprintf("%d-%s", id, act)
	headers := make(map[string]string)
	headers["producer"] = "mini-repos"
	headers["expires"] = strconv.FormatInt(time.Now().Unix()+expiresTime, 10)
	head, _ := json.Marshal(headers)
	a := base64.URLEncoding.EncodeToString(head)
	b := base64.URLEncoding.EncodeToString([]byte(content))
	c, _ := RsaSign([]byte(fmt.Sprintf("%s.%s", a, b)))
	return fmt.Sprintf("%s.%s.%s", a, b, c)
}

// ValidToken 验证并返回token,返回是否验证成功、id、act字段
func ValidToken(token string) (bool, int, string) {
	data := strings.Split(token, ".")
	if len(data) != 3 {
		return false, 0, ""
	}
	a := data[0]
	b := data[1]
	c := data[2]
	err := RsaVerify([]byte(fmt.Sprintf("%s.%s", a, b)), c)
	if err != nil {
		return false, 0, ""
	}
	h64, _ := base64.URLEncoding.DecodeString(a)
	headers := make(map[string]string)
	json.Unmarshal(h64, &headers)
	expires, _ := strconv.ParseInt(headers["expires"], 10, 64)
	if expires < time.Now().Unix() {
		return false, 0, ""
	}
	b64, _ := base64.URLEncoding.DecodeString(b)
	body := strings.Split(string(b64), "-")
	if len(body) != 2 {
		return false, 0, ""
	}
	id, _ := strconv.Atoi(body[0])
	return true, id, body[1]
}

func RsaVerify(data []byte, base64Sig string) error {
	bytes, err := base64.URLEncoding.DecodeString(base64Sig)
	if err != nil {
		return err
	}
	hashInstance := crypto.MD5.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	return rsa.VerifyPKCS1v15(&PublicKey, crypto.MD5, hashed, bytes)
}
func RsaSign(data []byte) (string, error) {
	hashInstance := crypto.MD5.New()
	hashInstance.Write(data)
	hashed := hashInstance.Sum(nil)
	bytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
func rsaGenerate(bits int) (*rsa.PrivateKey, rsa.PublicKey) {
	var privateKey *rsa.PrivateKey
	var publicKey rsa.PublicKey
	if exist, _ := CheckFileExist(KeyDir + "/privateKey.pem"); !exist {
		privateKey = createRsaPrivateKeyFile(bits)
		publicKey = createRsaPublicKeyFile(privateKey)
	} else {
		privateKey = readRsaPrivateKey()
		if exist, _ := CheckFileExist(KeyDir + "/publicKey.pem"); !exist {
			publicKey = createRsaPublicKeyFile(privateKey)
		} else {
			publicKey = *readRsaPublicKey()
		}
	}
	return privateKey, publicKey

}
func createRsaPrivateKeyFile(bits int) *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		Log.Fatal(err)
	}
	x509privateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyFile, err := os.Create(KeyDir + "/privateKey.pem")
	if err != nil {
		Log.Fatal(err)
	}
	defer privateKeyFile.Close()
	privateKeyBlock := pem.Block{
		Type:    "RSA Private Key",
		Headers: nil,
		Bytes:   x509privateKey,
	}
	pem.Encode(privateKeyFile, &privateKeyBlock)
	return privateKey
}
func createRsaPublicKeyFile(privateKey *rsa.PrivateKey) rsa.PublicKey {
	publicKey := privateKey.PublicKey
	x509publickey, _ := x509.MarshalPKIXPublicKey(&publicKey)
	publicKeyfile, err := os.Create(KeyDir + "/publicKey.pem")
	if err != nil {
		Log.Fatal(err)
	}
	defer publicKeyfile.Close()
	publicKeyBlock := pem.Block{
		Type:    "RSA Public Key",
		Headers: nil,
		Bytes:   x509publickey,
	}
	pem.Encode(publicKeyfile, &publicKeyBlock)
	return publicKey
}
func readRsaPublicKey() *rsa.PublicKey {
	file, err := os.Open(KeyDir + "/publicKey.pem")
	if err != nil {
		Log.Fatal(err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	block, _ := pem.Decode(buf)
	pki, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		Log.Fatal(err)
	}
	return pki.(*rsa.PublicKey)
}
func readRsaPrivateKey() *rsa.PrivateKey {
	file, err := os.Open(KeyDir + "/privateKey.pem")
	if err != nil {
		Log.Fatal(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem decode
	block, _ := pem.Decode(buf)
	//X509 decode
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		Log.Fatal(err)
	}
	return privateKey
}
