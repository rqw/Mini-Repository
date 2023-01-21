var x=Object.defineProperty,O=Object.defineProperties;var T=Object.getOwnPropertyDescriptors;var $=Object.getOwnPropertySymbols;var V=Object.prototype.hasOwnProperty,G=Object.prototype.propertyIsEnumerable;var M=(e,a,r)=>a in e?x(e,a,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[a]=r,_=(e,a)=>{for(var r in a||(a={}))V.call(a,r)&&M(e,r,a[r]);if($)for(var r of $(a))G.call(a,r)&&M(e,r,a[r]);return e},C=(e,a)=>O(e,T(a));var P=(e,a,r)=>new Promise((p,d)=>{var h=i=>{try{m(r.next(i))}catch(s){d(s)}},f=i=>{try{m(r.throw(i))}catch(s){d(s)}},m=i=>i.done?p(i.value):Promise.resolve(i.value).then(h,f);m((r=r.apply(e,a)).next())});import{R as L,a as j,bd as q,O as z,w as H,bk as J,b as K,bK as Q,aX as X,Q as F,o as g,h as w,n as U,H as E,j as S,l as W,t as A,J as Y,q as Z,c as ee,bm as ae,bG as te,bi as ne,dh as re,cD as oe,S as B}from"./index.f6ad53fc.js";import{B as D}from"./index.61a6744f.js";const ce=j({name:"LayoutBreadcrumb",components:{Icon:q,[D.name]:D},props:{theme:z.oneOf(["dark","light"])},setup(){const e=H([]),{currentRoute:a}=J(),{prefixCls:r}=K("layout-breadcrumb"),{getShowBreadCrumbIcon:p}=Q(),d=X(),{t:h}=ee();F(()=>P(this,null,function*(){var y,I,R;if(a.value.name===ae)return;const o=yield te(),t=a.value.matched,n=t==null?void 0:t[t.length-1];let c=a.value.path;n&&((y=n==null?void 0:n.meta)==null?void 0:y.currentActiveMenu)&&(c=n.meta.currentActiveMenu);const l=ne(o,c),b=o.filter(N=>N.path===l[0]),u=f(b,l);if(!u||u.length===0)return;const k=m(u);(I=a.value.meta)!=null&&I.currentActiveMenu&&k.push(C(_({},a.value),{name:((R=a.value.meta)==null?void 0:R.title)||a.value.name})),e.value=k}));function f(o,t){const n=[];return o.forEach(c=>{var l,b;t.includes(c.path)&&n.push(C(_({},c),{name:((l=c.meta)==null?void 0:l.title)||c.name})),(b=c.children)!=null&&b.length&&n.push(...f(c.children,t))}),n}function m(o){return re(o,t=>{const{meta:n,name:c}=t;if(!n)return!!c;const{title:l,hideBreadcrumb:b,hideMenu:u}=n;return!(!l||b||u)}).filter(t=>{var n;return!((n=t.meta)!=null&&n.hideBreadcrumb)})}function i(o,t,n){n==null||n.preventDefault();const{children:c,redirect:l,meta:b}=o;if((c==null?void 0:c.length)&&!l){n==null||n.stopPropagation();return}if(!(b!=null&&b.carryParam))if(l&&oe(l))d(l);else{let u="";t.length===1?u=t[0]:u=`${t.slice(1).pop()||""}`,u=/^\//.test(u)?u:`/${u}`,d(u)}}function s(o,t){return o.indexOf(t)!==o.length-1}function v(o){var t;return o.icon||((t=o.meta)==null?void 0:t.icon)}return{routes:e,t:h,prefixCls:r,getIcon:v,getShowBreadCrumbIcon:p,handleClick:i,hasRedirect:s}}}),se={key:1};function ue(e,a,r,p,d,h){const f=B("Icon"),m=B("router-link"),i=B("a-breadcrumb");return g(),w("div",{class:Z([e.prefixCls,`${e.prefixCls}--${e.theme}`])},[U(i,{routes:e.routes},{itemRender:E(({route:s,routes:v,paths:o})=>[e.getShowBreadCrumbIcon&&e.getIcon(s)?(g(),S(f,{key:0,icon:e.getIcon(s)},null,8,["icon"])):W("",!0),e.hasRedirect(v,s)?(g(),S(m,{key:2,to:"",onClick:t=>e.handleClick(s,o,t)},{default:E(()=>[Y(A(e.t(s.name||s.meta.title)),1)]),_:2},1032,["onClick"])):(g(),w("span",se,A(e.t(s.name||s.meta.title)),1))]),_:1},8,["routes"])],2)}var me=L(ce,[["render",ue]]);export{me as default};
