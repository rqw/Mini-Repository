import{R as b,a as k,bE as h,b as R,bU as v,C as O,bO as M,o as T,h as x,n as o,H as c,J as i,t as d,q as A,c as B,dp as N,k as C,c6 as f,d5 as $,d6 as F,S as p,N as P}from"./index.f6ad53fc.js";import{C as j}from"./CopyOutlined.f2dbc159.js";import{R as w}from"./index.0857c164.js";import"./useContentViewHeight.d656d165.js";import"./index.83733f95.js";const D=k({name:"SettingFooter",components:{CopyOutlined:j,RedoOutlined:w},setup(){const e=h(),{prefixCls:u}=R("setting-footer"),{t:s}=B(),{createSuccessModal:m,createMessage:r}=P(),g=v(),l=O(),t=M();function a(){const{isSuccessRef:n}=N(JSON.stringify(C(t.getProjectConfig),null,2));C(n)&&m({title:s("layout.setting.operatingTitle"),content:s("layout.setting.operatingContent")})}function S(){try{t.setProjectConfig(f);const{colorWeak:n,grayMode:_}=f;$(n),F(_),r.success(s("layout.setting.resetSuccess"))}catch(n){r.error(n)}}function y(){localStorage.clear(),t.resetAllState(),e.resetState(),g.resetState(),l.resetState(),location.reload()}return{prefixCls:u,t:s,handleCopy:a,handleResetSetting:S,handleClearAndRedo:y}}});function E(e,u,s,m,r,g){const l=p("CopyOutlined"),t=p("a-button"),a=p("RedoOutlined");return T(),x("div",{class:A(e.prefixCls)},[o(t,{type:"primary",block:"",onClick:e.handleCopy},{default:c(()=>[o(l,{class:"mr-2"}),i(" "+d(e.t("layout.setting.copyBtn")),1)]),_:1},8,["onClick"]),o(t,{color:"warning",block:"",onClick:e.handleResetSetting,class:"my-3"},{default:c(()=>[o(a,{class:"mr-2"}),i(" "+d(e.t("common.resetText")),1)]),_:1},8,["onClick"]),o(t,{color:"error",block:"",onClick:e.handleClearAndRedo},{default:c(()=>[o(a,{class:"mr-2"}),i(" "+d(e.t("layout.setting.clearBtn")),1)]),_:1},8,["onClick"])],2)}var q=b(D,[["render",E],["__scopeId","data-v-beabd496"]]);export{q as default};