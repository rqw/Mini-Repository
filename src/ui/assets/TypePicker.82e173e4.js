import{R as o,a as p,br as c,b as s,o as i,h as n,M as _,U as m,q as r,S as u,j as l,H as v,i as a,cW as b,cX as f}from"./index.f6ad53fc.js";const g=p({name:"MenuTypePicker",components:{Tooltip:c},props:{menuTypeList:{type:Array,defualt:()=>[]},handler:{type:Function,default:()=>({})},def:{type:String,default:""}},setup(){const{prefixCls:e}=s("setting-menu-type-picker");return{prefixCls:e}}}),h=e=>(b("data-v-91c9de94"),e=e(),f(),e),k=["onClick"],y=h(()=>a("div",{class:"mix-sidebar"},null,-1)),x=[y];function w(e,C,$,T,S,z){const d=u("Tooltip");return i(),n("div",{class:r(e.prefixCls)},[(i(!0),n(_,null,m(e.menuTypeList||[],t=>(i(),l(d,{key:t.title,title:t.title,placement:"bottom"},{default:v(()=>[a("div",{onClick:B=>e.handler(t),class:r([`${e.prefixCls}__item`,`${e.prefixCls}__item--${t.type}`,{[`${e.prefixCls}__item--active`]:e.def===t.type}])},x,10,k)]),_:2},1032,["title"]))),128))],2)}var L=o(g,[["render",w],["__scopeId","data-v-91c9de94"]]);export{L as default};