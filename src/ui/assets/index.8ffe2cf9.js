var ye=Object.defineProperty,ge=Object.defineProperties;var he=Object.getOwnPropertyDescriptors;var ie=Object.getOwnPropertySymbols;var xe=Object.prototype.hasOwnProperty,we=Object.prototype.propertyIsEnumerable;var se=(t,e,n)=>e in t?ye(t,e,{enumerable:!0,configurable:!0,writable:!0,value:n}):t[e]=n,E=(t,e)=>{for(var n in e||(e={}))xe.call(e,n)&&se(t,n,e[n]);if(ie)for(var n of ie(e))we.call(e,n)&&se(t,n,e[n]);return t},G=(t,e)=>ge(t,he(e));import{n as s,z as $,a5 as Se,w as I,cm as Ce,cn as _e,co as Ee,_ as j,M as je,a4 as T,a as N,v as De,bw as Fe,aN as re,ab as me,a6 as Be,y as ke,bL as le,f as O,aO as ae,a8 as Pe,ci as Re,a7 as Oe,b as Ae,cp as $e,k as l,aR as W,cq as Ie,Y as ze,cr as Le,cs as Te,bj as Ne,aT as Me,ct as Ue,Z as Ve,o as qe,j as Ge,H as ce,i as H,t as We,J as He,b$ as de}from"./index.f6ad53fc.js";import{T as Je}from"./index.6d5a7b38.js";import{P as Ye}from"./index.1719a346.js";import{g as Ze}from"./get.fb429fa9.js";import{G as ue,D as Ke,S as Qe}from"./siteSetting.c485f07c.js";import"./index.61a6744f.js";import"./index.b69ab639.js";import"./useSize.36511d9e.js";import"./eagerComputed.fb04c3da.js";import"./useContentViewHeight.d656d165.js";import"./ArrowLeftOutlined.8f7a54fe.js";import"./transButton.89c5992c.js";function L(t){return t!=null}var Xe=function(e){var n=e.itemPrefixCls,o=e.component,a=e.span,i=e.labelStyle,d=e.contentStyle,f=e.bordered,u=e.label,p=e.content,c=e.colon,_=o;if(f){var b;return s(_,{class:[(b={},$(b,"".concat(n,"-item-label"),L(u)),$(b,"".concat(n,"-item-content"),L(p)),b)],colSpan:a},{default:function(){return[L(u)&&s("span",{style:i},[u]),L(p)&&s("span",{style:d},[p])]}})}return s(_,{class:["".concat(n,"-item")],colSpan:a},{default:function(){return[s("div",{class:"".concat(n,"-item-container")},[u&&s("span",{class:["".concat(n,"-item-label"),$({},"".concat(n,"-item-no-colon"),!c)],style:i},[u]),p&&s("span",{class:"".concat(n,"-item-content"),style:d},[p])])]}})},J=Xe,et=function(e){var n=function(b,h,v){var D=h.colon,m=h.prefixCls,r=h.bordered,x=v.component,y=v.type,C=v.showLabel,B=v.showContent,k=v.labelStyle,F=v.contentStyle;return b.map(function(w,P){var R,S,g=w.props||{},z=g.prefixCls,M=z===void 0?m:z,K=g.span,Q=K===void 0?1:K,X=g.labelStyle,ee=g.contentStyle,te=g.label,ne=te===void 0?(S=(R=w.children)===null||R===void 0?void 0:R.label)===null||S===void 0?void 0:S.call(R):te,oe=Ce(w),U=_e(w),V=Ee(w),q=w.key;return typeof x=="string"?s(J,{key:"".concat(y,"-").concat(String(q)||P),class:U,style:V,labelStyle:j(j({},k.value),X),contentStyle:j(j({},F.value),ee),span:Q,colon:D,component:x,itemPrefixCls:M,bordered:r,label:C?ne:null,content:B?oe:null},null):[s(J,{key:"label-".concat(String(q)||P),class:U,style:j(j(j({},k.value),V),X),span:1,colon:D,component:x[0],itemPrefixCls:M,bordered:r,label:ne},null),s(J,{key:"content-".concat(String(q)||P),class:U,style:j(j(j({},F.value),V),ee),span:Q*2-1,component:x[1],itemPrefixCls:M,bordered:r,content:oe},null)]})},o=e.prefixCls,a=e.vertical,i=e.row,d=e.index,f=e.bordered,u=Se(be,{labelStyle:I({}),contentStyle:I({})}),p=u.labelStyle,c=u.contentStyle;return a?s(je,null,[s("tr",{key:"label-".concat(d),class:"".concat(o,"-row")},[n(i,e,{component:"th",type:"label",showLabel:!0,labelStyle:p,contentStyle:c})]),s("tr",{key:"content-".concat(d),class:"".concat(o,"-row")},[n(i,e,{component:"td",type:"content",showContent:!0,labelStyle:p,contentStyle:c})])]):s("tr",{key:d,class:"".concat(o,"-row")},[n(i,e,{component:f?["th","td"]:"td",type:"item",showLabel:!0,showContent:!0,labelStyle:p,contentStyle:c})])},tt=et;T.any;var nt=function(){return{prefixCls:String,label:T.any,labelStyle:{type:Object,default:void 0},contentStyle:{type:Object,default:void 0},span:{type:Number,default:1}}},ot=N({name:"ADescriptionsItem",props:nt(),slots:["label"],setup:function(e,n){var o=n.slots;return function(){var a;return(a=o.default)===null||a===void 0?void 0:a.call(o)}}}),ve={xxxl:3,xxl:3,xl:3,lg:3,md:3,sm:2,xs:1};function it(t,e){if(typeof t=="number")return t;if(me(t)==="object")for(var n=0;n<ae.length;n++){var o=ae[n];if(e[o]&&t[o]!==void 0)return t[o]||ve[o]}return 3}function pe(t,e,n){var o=t;return(e===void 0||e>n)&&(o=Re(t,{span:n}),Oe(e===void 0,"Descriptions","Sum of column `span` in a line not match `column` of Descriptions.")),o}function st(t,e){var n=Pe(t),o=[],a=[],i=e;return n.forEach(function(d,f){var u,p=(u=d.props)===null||u===void 0?void 0:u.span,c=p||1;if(f===n.length-1){a.push(pe(d,p,i)),o.push(a);return}c<i?(i-=c,a.push(d)):(a.push(pe(d,c,i)),o.push(a),i=e,a=[])}),o}var rt=function(){return{prefixCls:String,bordered:{type:Boolean,default:void 0},size:{type:String,default:"default"},title:T.any,extra:T.any,column:{type:[Number,Object],default:function(){return ve}},layout:String,colon:{type:Boolean,default:void 0},labelStyle:{type:Object,default:void 0},contentStyle:{type:Object,default:void 0}}},be=Symbol("descriptionsContext"),A=N({name:"ADescriptions",props:rt(),slots:["title","extra"],Item:ot,setup:function(e,n){var o=n.slots,a=De("descriptions",e),i=a.prefixCls,d=a.direction,f,u=I({});Fe(function(){f=re.subscribe(function(c){me(e.column)==="object"&&(u.value=c)})}),Be(function(){re.unsubscribe(f)}),ke(be,{labelStyle:le(e,"labelStyle"),contentStyle:le(e,"contentStyle")});var p=O(function(){return it(e.column,u.value)});return function(){var c,_,b,h,v=e.size,D=e.bordered,m=D===void 0?!1:D,r=e.layout,x=r===void 0?"horizontal":r,y=e.colon,C=y===void 0?!0:y,B=e.title,k=B===void 0?(_=o.title)===null||_===void 0?void 0:_.call(o):B,F=e.extra,w=F===void 0?(b=o.extra)===null||b===void 0?void 0:b.call(o):F,P=(h=o.default)===null||h===void 0?void 0:h.call(o),R=st(P,p.value);return s("div",{class:[i.value,(c={},$(c,"".concat(i.value,"-").concat(v),v!=="default"),$(c,"".concat(i.value,"-bordered"),!!m),$(c,"".concat(i.value,"-rtl"),d.value==="rtl"),c)]},[(k||w)&&s("div",{class:"".concat(i.value,"-header")},[k&&s("div",{class:"".concat(i.value,"-title")},[k]),w&&s("div",{class:"".concat(i.value,"-extra")},[w])]),s("div",{class:"".concat(i.value,"-view")},[s("table",null,[s("tbody",null,[R.map(function(S,g){return s(tt,{key:g,index:g,colon:C,prefixCls:i.value,vertical:x==="vertical",bordered:m,row:S},null)})])])])])}}});A.install=function(t){return t.component(A.name,A),t.component(A.Item.name,A.Item),t};var fe=A;function lt(t){return typeof t=="function"||Object.prototype.toString.call(t)==="[object Object]"&&!Ie(t)}const at={useCollapse:{type:Boolean,default:!0},title:{type:String,default:""},size:{type:String,validator:t=>["small","default","middle",void 0].includes(t),default:"small"},bordered:{type:Boolean,default:!0},column:{type:[Number,Object],default:()=>({xxl:4,xl:3,lg:3,md:3,sm:2,xs:1})},collapseOptions:{type:Object,default:null},schema:{type:Array,default:()=>[]},data:{type:Object}};var ct=N({name:"Description",props:at,emits:["register"],setup(t,{slots:e,emit:n}){const o=I(null),{prefixCls:a}=Ae("description"),i=$e(),d=O(()=>E(E({},t),l(o))),f=O(()=>G(E({},l(d)),{title:void 0})),u=O(()=>!!l(d).title),p=O(()=>E({canExpand:!1},l(f).collapseOptions)),c=O(()=>E(E({},l(i)),l(f)));function _(r){o.value=E(E({},l(o)),r)}function b({label:r,labelMinWidth:x,labelStyle:y}){if(!y&&!x)return r;const C=G(E({},y),{minWidth:`${x}px `});return s("div",{style:C},[r])}function h(){const{schema:r,data:x}=l(f);return l(r).map(y=>{const{render:C,field:B,span:k,show:F,contentMinWidth:w}=y;if(F&&W(F)&&!F(x))return null;const P=()=>{var z;const S=(z=l(f))==null?void 0:z.data;if(!S)return null;const g=Ze(S,B);return g&&!Ne(S).hasOwnProperty(B)?W(C)?C("",S):"":W(C)?C(g,S):g!=null?g:""},R=w;return s(fe.Item,{label:b(y),key:B,span:k},{default:()=>{if(!w)return P();const S={minWidth:`${R}px`};return s("div",{style:S},[P()])}})}).filter(y=>!!y)}const v=()=>{let r;return s(fe,ze({class:`${a}`},l(c)),lt(r=h())?r:{default:()=>[r]})},D=()=>{const r=t.useCollapse?v():s("div",null,[v()]);if(!t.useCollapse)return r;const{canExpand:x,helpMessage:y}=l(p),{title:C}=l(d);return s(Te,{title:C,canExpan:x,helpMessage:y},{default:()=>r,action:()=>Le(e,"action")})};return n("register",{setDescProps:_}),()=>l(u)?D():v()}});function Y(t){if(!Me())throw new Error("useDescription() can only be used inside setup() or functional components!");const e=I(null),n=I(!1);function o(i){l(n)&&Ue()||(e.value=i,t&&i.setDescProps(t),n.value=!0)}return[o,{setDescProps:i=>{var d;(d=l(e))==null||d.setDescProps(i)}}]}const Z=Ve(ct),dt={class:"flex justify-between items-center"},ut={class:"flex-1"},pt=["href"],ft=He(" \u662F\u4E00\u4E2A\u57FA\u4E8EVue3.0\u3001Vite\u3001 Ant-Design-Vue \u3001TypeScript \u7684\u540E\u53F0\u89E3\u51B3\u65B9\u6848\uFF0C\u76EE\u6807\u662F\u4E3A\u4E2D\u5927\u578B\u9879\u76EE\u5F00\u53D1,\u63D0\u4F9B\u73B0\u6210\u7684\u5F00\u7BB1\u89E3\u51B3\u65B9\u6848\u53CA\u4E30\u5BCC\u7684\u793A\u4F8B,\u539F\u5219\u4E0A\u4E0D\u4F1A\u9650\u5236\u4EFB\u4F55\u4EE3\u7801\u7528\u4E8E\u5546\u7528\u3002 "),Dt=N({__name:"index",setup(t){const{pkg:e,lastBuildTime:n}={pkg:{dependencies:{"@ant-design/colors":"^6.0.0","@ant-design/icons-vue":"^6.1.0","@iconify/iconify":"^2.2.1","@logicflow/core":"^1.1.13","@logicflow/extension":"^1.1.13","@vue/runtime-core":"^3.2.33","@vue/shared":"^3.2.33","@vueuse/core":"^8.3.0","@vueuse/shared":"^8.3.0","@zxcvbn-ts/core":"^2.0.1","ant-design-vue":"^3.2.0",axios:"^0.26.1",codemirror:"^5.65.3",cropperjs:"^1.5.12","crypto-js":"^4.1.1",dayjs:"^1.11.1",echarts:"^5.3.2","intro.js":"^5.1.0","lodash-es":"^4.17.21",mockjs:"^1.1.0",nprogress:"^0.2.0","path-to-regexp":"^6.2.0",pinia:"2.0.12",qs:"^6.10.3","resize-observer-polyfill":"^1.5.1",showdown:"^2.1.0",sortablejs:"^1.15.0",tinymce:"^5.10.3",vditor:"^3.8.13",vue:"^3.2.33","vue-i18n":"^9.1.9","vue-json-pretty":"^2.0.6","vue-router":"^4.0.14","vue-types":"^4.1.1",xlsx:"^0.18.5"},devDependencies:{"@commitlint/cli":"^16.2.3","@commitlint/config-conventional":"^16.2.1","@iconify/json":"^2.1.30","@purge-icons/generated":"^0.8.1","@types/codemirror":"^5.60.5","@types/crypto-js":"^4.1.1","@types/fs-extra":"^9.0.13","@types/inquirer":"^8.2.1","@types/intro.js":"^3.0.2","@types/lodash-es":"^4.17.6","@types/mockjs":"^1.0.6","@types/node":"^17.0.25","@types/nprogress":"^0.2.0","@types/qs":"^6.9.7","@types/showdown":"^1.9.4","@types/sortablejs":"^1.10.7","@typescript-eslint/eslint-plugin":"^5.20.0","@typescript-eslint/parser":"^5.20.0","@vitejs/plugin-legacy":"^1.8.1","@vitejs/plugin-vue":"^2.3.1","@vitejs/plugin-vue-jsx":"^1.3.10","@vue/compiler-sfc":"^3.2.33","@vue/test-utils":"^2.0.0-rc.21",autoprefixer:"^10.4.4","conventional-changelog-cli":"^2.2.2","cross-env":"^7.0.3","cz-git":"^1.3.11",czg:"^1.3.11",dotenv:"^16.0.0",eslint:"^8.13.0","eslint-config-prettier":"^8.5.0","eslint-plugin-prettier":"^4.0.0","eslint-plugin-vue":"^8.6.0",esno:"^0.14.1","fs-extra":"^10.1.0",husky:"^7.0.4",inquirer:"^8.2.2",less:"^4.1.2","lint-staged":"12.3.7","npm-run-all":"^4.1.5",picocolors:"^1.0.0",postcss:"^8.4.12","postcss-html":"^1.4.1","postcss-less":"^6.0.0",prettier:"^2.6.2",rimraf:"^3.0.2",rollup:"^2.70.2","rollup-plugin-visualizer":"^5.6.0",stylelint:"^14.7.1","stylelint-config-prettier":"^9.0.3","stylelint-config-recommended":"^7.0.0","stylelint-config-recommended-vue":"^1.4.0","stylelint-config-standard":"^25.0.0","stylelint-order":"^5.0.0","ts-node":"^10.7.0",typescript:"^4.6.3",vite:"^2.9.5","vite-plugin-compression":"^0.5.1","vite-plugin-html":"^3.2.0","vite-plugin-imagemin":"^0.6.1","vite-plugin-mkcert":"^1.6.0","vite-plugin-mock":"^2.9.6","vite-plugin-purge-icons":"^0.8.1","vite-plugin-pwa":"^0.11.13","vite-plugin-style-import":"^2.0.0","vite-plugin-svg-icons":"^2.0.1","vite-plugin-theme":"^0.8.6","vite-plugin-vue-setup-extend":"^0.4.0","vite-plugin-windicss":"^1.8.4","vue-eslint-parser":"^8.3.0","vue-tsc":"^0.33.9"},name:"vben-admin",version:"2.8.0"},lastBuildTime:"2023-01-21 15:50:30"},{dependencies:o,devDependencies:a,name:i,version:d}=e,f=[],u=[],p=m=>r=>de(Je,{color:m},()=>r),c=m=>r=>de("a",{href:r,target:"_blank"},m),_=[{label:"\u7248\u672C",field:"version",render:p("blue")},{label:"\u6700\u540E\u7F16\u8BD1\u65F6\u95F4",field:"lastBuildTime",render:p("blue")},{label:"\u6587\u6863\u5730\u5740",field:"doc",render:c("\u6587\u6863\u5730\u5740")},{label:"\u9884\u89C8\u5730\u5740",field:"preview",render:c("\u9884\u89C8\u5730\u5740")},{label:"Github",field:"github",render:c("Github")}],b={version:d,lastBuildTime:n,doc:Ke,preview:Qe,github:ue};Object.keys(o).forEach(m=>{f.push({field:m,label:m})}),Object.keys(a).forEach(m=>{u.push({field:m,label:m})});const[h]=Y({title:"\u751F\u4EA7\u73AF\u5883\u4F9D\u8D56",data:o,schema:f,column:3}),[v]=Y({title:"\u5F00\u53D1\u73AF\u5883\u4F9D\u8D56",data:a,schema:u,column:3}),[D]=Y({title:"\u9879\u76EE\u4FE1\u606F",data:b,schema:_,column:2});return(m,r)=>(qe(),Ge(l(Ye),{title:"\u5173\u4E8E"},{headerContent:ce(()=>[H("div",dt,[H("span",ut,[H("a",{href:l(ue),target:"_blank"},We(l(i)),9,pt),ft])])]),default:ce(()=>[s(l(Z),{onRegister:l(D),class:"enter-y"},null,8,["onRegister"]),s(l(Z),{onRegister:l(h),class:"my-4 enter-y"},null,8,["onRegister"]),s(l(Z),{onRegister:l(v),class:"enter-y"},null,8,["onRegister"])]),_:1}))}});export{Dt as default};
