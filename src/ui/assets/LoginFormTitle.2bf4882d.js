var T=Object.defineProperty;var P=Object.getOwnPropertySymbols;var w=Object.prototype.hasOwnProperty,S=Object.prototype.propertyIsEnumerable;var h=(e,o,t)=>o in e?T(e,o,{enumerable:!0,configurable:!0,writable:!0,value:t}):e[o]=t,y=(e,o)=>{for(var t in o||(o={}))w.call(o,t)&&h(e,t,o[t]);if(P)for(var t of P(o))S.call(o,t)&&h(e,t,o[t]);return e};var R=(e,o,t)=>new Promise((r,i)=>{var g=n=>{try{a(t.next(n))}catch(s){i(s)}},_=n=>{try{a(t.throw(n))}catch(s){i(s)}},a=n=>n.done?r(n.value):Promise.resolve(n.value).then(g,_);a((t=t.apply(e,o)).next())});import{w as v,f as c,c as O,k as l,a as I,o as x,h as E,t as j}from"./index.f6ad53fc.js";var d=(e=>(e[e.LOGIN=0]="LOGIN",e[e.REGISTER=1]="REGISTER",e[e.RESET_PASSWORD=2]="RESET_PASSWORD",e[e.MOBILE=3]="MOBILE",e[e.QR_CODE=4]="QR_CODE",e))(d||{});const F=v(0);function B(){function e(r){F.value=r}const o=c(()=>F.value);function t(){e(0)}return{setLoginState:e,getLoginState:o,handleBackLogin:t}}function A(e){function o(){return R(this,null,function*(){const t=l(e);return t?yield t.validate():void 0})}return{validForm:o}}function C(e){const{t:o}=O(),t=c(()=>f(o("sys.login.accountPlaceholder"))),r=c(()=>f(o("sys.login.passwordPlaceholder"))),i=c(()=>f(o("sys.login.smsPlaceholder"))),g=c(()=>f(o("sys.login.mobilePlaceholder"))),_=(s,u)=>R(this,null,function*(){return u?Promise.resolve():Promise.reject(o("sys.login.policyPlaceholder"))}),a=s=>(u,m)=>R(this,null,function*(){return m?m!==s?Promise.reject(o("sys.login.diffPwd")):Promise.resolve():Promise.reject(o("sys.login.passwordPlaceholder"))});return{getFormRules:c(()=>{const s=l(t),u=l(r),m=l(i),b=l(g),p={sms:m,mobile:b};switch(l(F)){case 1:return y({account:s,password:u,confirmPassword:[{validator:a(e==null?void 0:e.password),trigger:"change"}],policy:[{validator:_,trigger:"change"}]},p);case 2:return y({account:s},p);case 3:return p;default:return{account:s,password:u}}})}}function f(e){return[{required:!0,message:e,trigger:"change"}]}const G={class:"mb-3 text-2xl font-bold text-center xl:text-3xl enter-x xl:text-left"},L=I({__name:"LoginFormTitle",setup(e){const{t:o}=O(),{getLoginState:t}=B(),r=c(()=>({[d.RESET_PASSWORD]:o("sys.login.forgetFormTitle"),[d.LOGIN]:o("sys.login.signInFormTitle"),[d.REGISTER]:o("sys.login.signUpFormTitle"),[d.MOBILE]:o("sys.login.mobileSignInFormTitle")})[l(t)]);return(i,g)=>(x(),E("h2",G,j(l(r)),1))}});var N=Object.freeze(Object.defineProperty({__proto__:null,default:L},Symbol.toStringTag,{value:"Module"}));export{d as L,L as _,C as a,A as b,N as c,B as u};
