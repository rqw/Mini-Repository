var P=(t,r,l)=>new Promise((f,b)=>{var o=s=>{try{p(l.next(s))}catch(m){b(m)}},x=s=>{try{p(l.throw(s))}catch(m){b(m)}},p=s=>s.done?f(s.value):Promise.resolve(s.value).then(o,x);p((l=l.apply(t,r)).next())});import{r as U,F as G,s as J,a as H,v as W,w as I,x as z,f as O,y as Q,n,z as X,B as Y,I as N,c as Z,b as ee,C as ae,D as te,k as e,o as ne,h as re,E as j,G as K,H as d,J as R,t as E,K as L,L as le,M as se,N as oe}from"./index.f6ad53fc.js";import{c as ue,C as ie,a as k}from"./index.2f09a5b0.js";import{F as C}from"./index.a47b0923.js";import"./index.62e961d2.js";import{u as ce,a as de,L as A,_ as fe,b as me}from"./LoginFormTitle.2bf4882d.js";import{C as ve,R as ge}from"./index.9b9d0162.js";import"./get.fb429fa9.js";import"./useSize.36511d9e.js";C.useInjectFormItemContext=U;C.ItemRest=G;C.install=function(t){return t.component(C.name,C),t.component(C.Item.name,C.Item),t.component(G.name,G),t};function pe(t,r){var l=typeof Symbol!="undefined"&&t[Symbol.iterator]||t["@@iterator"];if(!l){if(Array.isArray(t)||(l=J(t))||r&&t&&typeof t.length=="number"){l&&(t=l);var f=0,b=function(){};return{s:b,n:function(){return f>=t.length?{done:!0}:{done:!1,value:t[f++]}},e:function(m){throw m},f:b}}throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}var o=!0,x=!1,p;return{s:function(){l=l.call(t)},n:function(){var m=l.next();return o=m.done,m},e:function(m){x=!0,p=m},f:function(){try{!o&&l.return!=null&&l.return()}finally{if(x)throw p}}}}var $=H({name:"ACheckboxGroup",props:ue(),setup:function(r,l){var f=l.slots,b=l.emit,o=l.expose,x=U(),p=W("checkbox",r),s=p.prefixCls,m=p.direction,y=I((r.value===void 0?r.defaultValue:r.value)||[]);z(function(){return r.value},function(){y.value=r.value||[]});var w=O(function(){return r.options.map(function(u){return typeof u=="string"||typeof u=="number"?{label:u,value:u}:u})}),S=I(Symbol()),_=I(new Map),F=function(i){_.value.delete(i),S.value=Symbol()},M=function(i,v){_.value.set(i,v),S.value=Symbol()},h=I(new Map);z(S,function(){var u=new Map,i=pe(_.value.values()),v;try{for(i.s();!(v=i.n()).done;){var a=v.value;u.set(a,!0)}}catch(c){i.e(c)}finally{i.f()}h.value=u});var T=function(i){var v=y.value.indexOf(i.value),a=Y(y.value);v===-1?a.push(i.value):a.splice(v,1),r.value===void 0&&(y.value=a);var c=a.filter(function(g){return h.value.has(g)}).sort(function(g,V){var D=w.value.findIndex(function(B){return B.value===g}),q=w.value.findIndex(function(B){return B.value===V});return D-q});b("update:value",c),b("change",c),x.onFieldChange()};return Q(ie,{cancelValue:F,registerValue:M,toggleOption:T,mergedValue:y,name:O(function(){return r.name}),disabled:O(function(){return r.disabled})}),o({mergedValue:y}),function(){var u,i=r.id,v=i===void 0?x.id.value:i,a=null,c="".concat(s.value,"-group");return w.value&&w.value.length>0&&(a=w.value.map(function(g){var V;return n(k,{prefixCls:s.value,key:g.value.toString(),disabled:"disabled"in g?g.disabled:r.disabled,indeterminate:g.indeterminate,value:g.value,checked:y.value.indexOf(g.value)!==-1,onChange:g.onChange,class:"".concat(c,"-item")},{default:function(){return[g.label===void 0?(V=f.label)===null||V===void 0?void 0:V.call(f,g):g.label]}})})),n("div",{class:[c,X({},"".concat(c,"-rtl"),m.value==="rtl")],id:v},[a||((u=f.default)===null||u===void 0?void 0:u.call(f))])}}});k.Group=$;k.install=function(t){return t.component(k.name,k),t.component($.name,$),t};const Se=H({__name:"LoginForm",setup(t){const r=ve,l=ge,f=C.Item,b=N.Password,{t:o}=Z(),{notification:x,createErrorModal:p}=oe(),{prefixCls:s}=ee("login"),m=ae(),{setLoginState:y,getLoginState:w}=ce(),{getFormRules:S}=de(),_=I(),F=I(!1),M=I(!1),h=te({account:"vben",password:"123456"}),{validForm:T}=me(_),u=O(()=>e(w)===A.LOGIN);function i(){return P(this,null,function*(){const v=yield T();if(!!v)try{F.value=!0;const a=yield m.login({password:v.password,username:v.account,mode:"none"});a&&x.success({message:o("sys.login.loginSuccessTitle"),description:`${o("sys.login.loginSuccessDesc")}: ${a.realName}`,duration:3})}catch(a){p({title:o("sys.api.errorTip"),content:a.message||o("sys.api.networkExceptionMsg"),getContainer:()=>document.body.querySelector(`.${s}`)||document.body})}finally{F.value=!1}})}return(v,a)=>(ne(),re(se,null,[j(n(fe,{class:"enter-x"},null,512),[[K,e(u)]]),j(n(e(C),{class:"p-4 enter-x",model:h,rules:e(S),ref_key:"formRef",ref:_,onKeypress:le(i,["enter"])},{default:d(()=>[n(e(f),{name:"account",class:"enter-x"},{default:d(()=>[n(e(N),{size:"large",value:h.account,"onUpdate:value":a[0]||(a[0]=c=>h.account=c),placeholder:e(o)("sys.login.userName"),class:"fix-auto-fill"},null,8,["value","placeholder"])]),_:1}),n(e(f),{name:"password",class:"enter-x"},{default:d(()=>[n(e(b),{size:"large",visibilityToggle:"",value:h.password,"onUpdate:value":a[1]||(a[1]=c=>h.password=c),placeholder:e(o)("sys.login.password")},null,8,["value","placeholder"])]),_:1}),n(e(l),{class:"enter-x"},{default:d(()=>[n(e(r),{span:12},{default:d(()=>[n(e(f),null,{default:d(()=>[n(e(k),{checked:M.value,"onUpdate:checked":a[2]||(a[2]=c=>M.value=c),size:"small"},{default:d(()=>[R(E(e(o)("sys.login.rememberMe")),1)]),_:1},8,["checked"])]),_:1})]),_:1}),n(e(r),{span:12},{default:d(()=>[n(e(f),{style:{"text-align":"right"}},{default:d(()=>[n(e(L),{type:"link",size:"small",onClick:a[3]||(a[3]=c=>e(y)(e(A).RESET_PASSWORD))},{default:d(()=>[R(E(e(o)("sys.login.forgetPassword")),1)]),_:1})]),_:1})]),_:1})]),_:1}),n(e(f),{class:"enter-x"},{default:d(()=>[n(e(L),{type:"primary",size:"large",block:"",onClick:i,loading:F.value},{default:d(()=>[R(E(e(o)("sys.login.loginButton")),1)]),_:1},8,["loading"])]),_:1}),n(e(l),{class:"enter-x"},{default:d(()=>[n(e(r),{md:8,xs:24},{default:d(()=>[n(e(L),{block:"",onClick:a[4]||(a[4]=c=>e(y)(e(A).MOBILE))},{default:d(()=>[R(E(e(o)("sys.login.mobileSignInFormTitle")),1)]),_:1})]),_:1}),n(e(r),{md:6,xs:24},{default:d(()=>[n(e(L),{block:"",onClick:a[5]||(a[5]=c=>e(y)(e(A).REGISTER))},{default:d(()=>[R(E(e(o)("sys.login.registerButton")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model","rules","onKeypress"]),[[K,e(u)]])],64))}});export{Se as default};
