var H=(t,i,e)=>new Promise((n,a)=>{var r=o=>{try{c(e.next(o))}catch(d){a(d)}},s=o=>{try{c(e.throw(o))}catch(d){a(d)}},c=o=>o.done?n(o.value):Promise.resolve(o.value).then(r,s);c((e=e.apply(t,i)).next())});import{bo as w,c8 as m,aQ as l,bA as v,w as u,f as y,k as g}from"./index.f6ad53fc.js";const b=Symbol();function z(t){return w(t,b,{native:!0})}function p(t,i=150,e){let n=()=>{t()};n=v(n,i);const r=()=>{e&&e.immediate&&n(),window.addEventListener("resize",n)},s=()=>{window.removeEventListener("resize",n)};return m(()=>{r()}),l(()=>{s()}),[r,s]}const h=u(0),f=u(0);function C(){function t(e){h.value=e}function i(e){f.value=e}return{headerHeightRef:h,footerHeightRef:f,setHeaderHeight:t,setFooterHeight:i}}function F(){const t=u(window.innerHeight),i=u(window.innerHeight),e=y(()=>g(t)-g(h)-g(f)||0);p(()=>{t.value=window.innerHeight},100,{immediate:!0});function n(a){return H(this,null,function*(){i.value=a})}z({contentHeight:e,setPageHeight:n,pageHeight:i})}export{C as a,p as b,F as u};