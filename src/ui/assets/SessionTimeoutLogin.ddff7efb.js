import{R as a,a as u,b as m,C as p,bE as c,bO as d,w as l,a0 as f,a6 as _,cQ as g,o as v,j as b,H as S,i as x,n as B,q as C,k as I,bY as M}from"./index.f6ad53fc.js";import T from"./Login.ab908ad9.js";import"./LoginForm.0aed27fa.js";import"./index.2f09a5b0.js";import"./index.a47b0923.js";import"./index.62e961d2.js";import"./get.fb429fa9.js";import"./useSize.36511d9e.js";import"./LoginFormTitle.2bf4882d.js";import"./index.9b9d0162.js";import"./ForgetPasswordForm.dc719a5f.js";import"./index.20b37200.js";import"./RegisterForm.d945a584.js";import"./MobileForm.6f99ac28.js";const h=u({__name:"SessionTimeoutLogin",setup(k){const{prefixCls:t}=m("st-login"),e=p(),i=c(),n=d(),o=l(0),r=()=>n.getProjectConfig.permissionMode===g.BACK;return f(()=>{var s;o.value=(s=e.getUserInfo)==null?void 0:s.userId,e.getUserInfo}),_(()=>{(o.value&&o.value!==e.getUserInfo.userId||r()&&i.getLastBuildMenuTime===0)&&document.location.reload()}),(s,U)=>(v(),b(M,null,{default:S(()=>[x("div",{class:C(I(t))},[B(T,{sessionTimeout:""})],2)]),_:1}))}});var O=a(h,[["__scopeId","data-v-2f8beb6e"]]);export{O as default};