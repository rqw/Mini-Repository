var be=Object.defineProperty,me=Object.defineProperties;var xe=Object.getOwnPropertyDescriptors;var ee=Object.getOwnPropertySymbols;var we=Object.prototype.hasOwnProperty,ye=Object.prototype.propertyIsEnumerable;var te=(e,t,r)=>t in e?be(e,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[t]=r,W=(e,t)=>{for(var r in t||(t={}))we.call(t,r)&&te(e,r,t[r]);if(ee)for(var r of ee(t))ye.call(t,r)&&te(e,r,t[r]);return e},ae=(e,t)=>me(e,xe(t));var ne=(e,t,r)=>new Promise((v,n)=>{var x=f=>{try{p(r.next(f))}catch(w){n(w)}},g=f=>{try{p(r.throw(f))}catch(w){n(w)}},p=f=>f.done?v(f.value):Promise.resolve(f.value).then(x,g);p((r=r.apply(e,t)).next())});import{n as h,aV as _e,aP as He,a as Y,a4 as R,v as Ce,w as T,cx as ke,f as C,a8 as Pe,a2 as $e,z as Z,cy as Fe,cg as Be,aE as Re,cz as Se,b as de,bf as Te,R as ce,o as j,h as q,i as J,V as S,q as N,bt as ge,a0 as ze,a1 as V,cA as Ae,k as $,cB as Oe,cC as Ie,x as ue,cD as Ee,cE as je,O as _,y as Ne,aI as re,S as oe,j as ie,T as We,U as De,H as D,W as Ve,X as Le,M as Me,J as Ue,t as Xe,Y as Ze,l as le,Z as pe}from"./index.f6ad53fc.js";import{B as qe}from"./index.61a6744f.js";import{A as Je}from"./index.b69ab639.js";import{b as Ye,a as Ge}from"./useContentViewHeight.d656d165.js";import{A as Qe}from"./ArrowLeftOutlined.8f7a54fe.js";import{T as Ke}from"./transButton.89c5992c.js";var et={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M869 487.8L491.2 159.9c-2.9-2.5-6.6-3.9-10.5-3.9h-88.5c-7.4 0-10.8 9.2-5.2 14l350.2 304H152c-4.4 0-8 3.6-8 8v60c0 4.4 3.6 8 8 8h585.1L386.9 854c-5.6 4.9-2.2 14 5.2 14h91.5c1.9 0 3.8-.7 5.2-2L869 536.2a32.07 32.07 0 000-48.4z"}}]},name:"arrow-right",theme:"outlined"},tt=et;function se(e){for(var t=1;t<arguments.length;t++){var r=arguments[t]!=null?Object(arguments[t]):{},v=Object.keys(r);typeof Object.getOwnPropertySymbols=="function"&&(v=v.concat(Object.getOwnPropertySymbols(r).filter(function(n){return Object.getOwnPropertyDescriptor(r,n).enumerable}))),v.forEach(function(n){at(e,n,r[n])})}return e}function at(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}var G=function(t,r){var v=se({},t,r.attrs);return h(_e,se({},v,{icon:tt}),null)};G.displayName="ArrowRightOutlined";G.inheritAttrs=!1;var nt=G,rt=function(){return{backIcon:R.any,prefixCls:String,title:R.any,subTitle:R.any,breadcrumb:R.object,tags:R.any,footer:R.any,extra:R.any,avatar:R.object,ghost:{type:Boolean,default:void 0},onBack:Function}},ot=Y({name:"APageHeader",props:rt(),slots:["backIcon","avatar","breadcrumb","title","subTitle","tags","extra","footer"],setup:function(t,r){var v=r.emit,n=r.slots,x=Ce("page-header",t),g=x.prefixCls,p=x.direction,f=x.pageHeader,w=T(!1),O=ke(),k=function(a){var l=a.width;O.value||(w.value=l<768)},F=C(function(){var i,a,l;return(l=(i=t.ghost)!==null&&i!==void 0?i:(a=f.value)===null||a===void 0?void 0:a.ghost)!==null&&l!==void 0?l:!0}),B=function(){var a,l,c;return(c=(a=t.backIcon)!==null&&a!==void 0?a:(l=n.backIcon)===null||l===void 0?void 0:l.call(n))!==null&&c!==void 0?c:p.value==="rtl"?h(nt,null,null):h(Qe,null,null)},u=function(a){return!a||!t.onBack?null:h(Se,{componentName:"PageHeader",children:function(c){var o=c.back;return h("div",{class:"".concat(g.value,"-back")},[h(Ke,{onClick:function(d){v("back",d)},class:"".concat(g.value,"-back-button"),"aria-label":o},{default:function(){return[a]}})])}},null)},P=function(){var a;return t.breadcrumb?h(qe,t.breadcrumb,null):(a=n.breadcrumb)===null||a===void 0?void 0:a.call(n)},y=function(){var a,l,c,o,s,d,H,z,L,M=t.avatar,I=(a=t.title)!==null&&a!==void 0?a:(l=n.title)===null||l===void 0?void 0:l.call(n),E=(c=t.subTitle)!==null&&c!==void 0?c:(o=n.subTitle)===null||o===void 0?void 0:o.call(n),U=(s=t.tags)!==null&&s!==void 0?s:(d=n.tags)===null||d===void 0?void 0:d.call(n),X=(H=t.extra)!==null&&H!==void 0?H:(z=n.extra)===null||z===void 0?void 0:z.call(n),A="".concat(g.value,"-heading"),Q=I||E||U||X;if(!Q)return null;var he=B(),K=u(he),ve=K||M||Q;return h("div",{class:A},[ve&&h("div",{class:"".concat(A,"-left")},[K,M?h(Je,M,null):(L=n.avatar)===null||L===void 0?void 0:L.call(n),I&&h("span",{class:"".concat(A,"-title"),title:typeof I=="string"?I:void 0},[I]),E&&h("span",{class:"".concat(A,"-sub-title"),title:typeof E=="string"?E:void 0},[E]),U&&h("span",{class:"".concat(A,"-tags")},[U])]),X&&h("span",{class:"".concat(A,"-extra")},[X])])},b=function(){var a,l,c=(a=t.footer)!==null&&a!==void 0?a:Re((l=n.footer)===null||l===void 0?void 0:l.call(n));return Fe(c)?null:h("div",{class:"".concat(g.value,"-footer")},[c])},m=function(a){return h("div",{class:"".concat(g.value,"-content")},[a])};return function(){var i,a,l,c=((a=t.breadcrumb)===null||a===void 0?void 0:a.routes)||n.breadcrumb,o=t.footer||n.footer,s=Pe((l=n.default)===null||l===void 0?void 0:l.call(n)),d=$e(g.value,(i={"has-breadcrumb":c,"has-footer":o},Z(i,"".concat(g.value,"-ghost"),F.value),Z(i,"".concat(g.value,"-rtl"),p.value==="rtl"),Z(i,"".concat(g.value,"-compact"),w.value),i));return h(Be,{onResize:k},{default:function(){return[h("div",{class:d},[P(),y(),s.length?m(s):null,b()])]}})}}}),it=He(ot);const lt=Y({name:"PageFooter",inheritAttrs:!1,setup(){const{prefixCls:e}=de("page-footer"),{getCalcContentWidth:t}=Te();return{prefixCls:e,getCalcContentWidth:t}}});function st(e,t,r,v,n,x){return j(),q("div",{class:N(e.prefixCls),style:ge({width:e.getCalcContentWidth})},[J("div",{class:N(`${e.prefixCls}__left`)},[S(e.$slots,"left",{},void 0,!0)],2),S(e.$slots,"default",{},void 0,!0),J("div",{class:N(`${e.prefixCls}__right`)},[S(e.$slots,"right",{},void 0,!0)],2)],6)}var fe=ce(lt,[["render",st],["__scopeId","data-v-39c0958e"]]);function dt(e){let t;ze(()=>{e(),V(()=>{t=!0})}),Ae(()=>{t&&e()})}function ct(e,t,r,v,n=0,x=T(0)){const g=T(null),{footerHeightRef:p}=Ge();let f={useLayoutFooter:!0};const w=u=>{f=u};function O(){V(()=>{B()})}function k(u,P="all"){var i,a,l,c;function y(o){return Number(o.replace(/[^\d]/g,""))}let b=0;const m="0px";if(u){const o=getComputedStyle(u),s=y((i=o==null?void 0:o.marginTop)!=null?i:m),d=y((a=o==null?void 0:o.marginBottom)!=null?a:m),H=y((l=o==null?void 0:o.paddingTop)!=null?l:m),z=y((c=o==null?void 0:o.paddingBottom)!=null?c:m);P==="all"?(b+=s,b+=d,b+=H,b+=z):P==="top"?(b+=s,b+=H):(b+=d,b+=z)}return b}function F(u){return u==null?null:u instanceof HTMLDivElement?u:u.$el}function B(){return ne(this,null,function*(){var c;if(!e.value)return;yield V();const u=F($(t));if(!u)return;const{bottomIncludeBody:P}=Oe(u);let y=0;r.forEach(o=>{var s,d;y+=(d=(s=F($(o)))==null?void 0:s.offsetHeight)!=null?d:0});let b=(c=k(u))!=null?c:0;v.forEach(o=>{b+=k(F($(o)))});let m=0;function i(o,s){if(o&&s){const d=o.parentElement;d&&(Ee(s)?d.classList.contains(s)?m+=k(d,"bottom"):(m+=k(d,"bottom"),i(d,s)):je(s)&&s>0&&(m+=k(d,"bottom"),i(d,--s)))}}Ie(n)?i(u,$(n)):i(u,n);let a=P-$(p)-$(x)-y-b-m;const l=()=>{var o;(o=f.elements)==null||o.forEach(s=>{var d,H;a+=(H=(d=F($(s)))==null?void 0:d.offsetHeight)!=null?H:0})};f.useLayoutFooter&&$(p)>0,l(),g.value=a})}return dt(()=>{V(()=>{B()})}),Ye(()=>{B()},50,{immediate:!0}),ue(()=>[p.value],()=>{B()},{flush:"post",immediate:!0}),{redoHeight:O,setCompensation:w,contentHeight:g}}const gt=Y({name:"PageWrapper",components:{PageFooter:fe,PageHeader:it},inheritAttrs:!1,props:{title:_.string,dense:_.bool,ghost:_.bool,content:_.string,contentStyle:{type:Object},contentBackground:_.bool,contentFullHeight:_.bool,contentClass:_.string,fixedHeight:_.bool,upwardSpace:_.oneOfType([_.number,_.string]).def(0)},setup(e,{slots:t,attrs:r}){const v=T(null),n=T(null),x=T(null),g=T(null),{prefixCls:p}=de("page-wrapper");Ne(ft,C(()=>e.fixedHeight));const f=C(()=>e.contentFullHeight),w=C(()=>e.upwardSpace),{redoHeight:O,setCompensation:k,contentHeight:F}=ct(f,v,[n,g],[x],w);k({useLayoutFooter:!0,elements:[g]});const B=C(()=>{var i;return[p,{[`${p}--dense`]:e.dense},(i=r.class)!=null?i:{}]}),u=C(()=>e.content||(t==null?void 0:t.headerContent)||e.title||y.value.length),P=C(()=>(t==null?void 0:t.leftFooter)||(t==null?void 0:t.rightFooter)),y=C(()=>Object.keys(re(t,"default","leftFooter","rightFooter","headerContent"))),b=C(()=>{const{contentFullHeight:i,contentStyle:a,fixedHeight:l}=e;if(!i)return W({},a);const c=`${$(F)}px`;return W(ae(W({},a),{minHeight:c}),l?{height:c}:{})}),m=C(()=>{const{contentBackground:i,contentClass:a}=e;return[`${p}-content`,a,{[`${p}-content-bg`]:i}]});return ue(()=>[P.value],()=>{O()},{flush:"post",immediate:!0}),{getContentStyle:b,wrapperRef:v,headerRef:n,contentRef:x,footerRef:g,getClass:B,getHeaderSlots:y,prefixCls:p,getShowHeader:u,getShowFooter:P,omit:re,getContentClass:m}}});function ut(e,t,r,v,n,x){const g=oe("PageHeader"),p=oe("PageFooter");return j(),q("div",{class:N(e.getClass),ref:"wrapperRef"},[e.getShowHeader?(j(),ie(g,Ze({key:0,ghost:e.ghost,title:e.title},e.omit(e.$attrs,"class"),{ref:"headerRef"}),We({default:D(()=>[e.content?(j(),q(Me,{key:0},[Ue(Xe(e.content),1)],64)):S(e.$slots,"headerContent",{key:1})]),_:2},[De(e.getHeaderSlots,f=>({name:f,fn:D(w=>[S(e.$slots,f,Ve(Le(w||{})))])}))]),1040,["ghost","title"])):le("",!0),J("div",{class:N(["overflow-hidden",e.getContentClass]),style:ge(e.getContentStyle),ref:"contentRef"},[S(e.$slots,"default")],6),e.getShowFooter?(j(),ie(p,{key:1,ref:"footerRef"},{left:D(()=>[S(e.$slots,"leftFooter")]),right:D(()=>[S(e.$slots,"rightFooter")]),_:3},512)):le("",!0)],2)}var pt=ce(gt,[["render",ut]]);pe(fe);const _t=pe(pt),ft="PageWrapperFixedHeight";export{_t as P};
