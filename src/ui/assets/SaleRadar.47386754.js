import{a as i,w as r,x as s,o as l,j as n,H as d,i as m,bt as u,k as f}from"./index.f6ad53fc.js";import{C as c}from"./index.2ace8746.js";import"./index.83733f95.js";import"./index.62e961d2.js";import{u as h}from"./useECharts.546bc9e3.js";import"./index.9b9d0162.js";const x=i({__name:"SaleRadar",props:{loading:Boolean,width:{type:String,default:"100%"},height:{type:String,default:"400px"}},setup(e){const a=e,t=r(null),{setOptions:o}=h(t);return s(()=>a.loading,()=>{a.loading||o({legend:{bottom:0,data:["Visits","Sales"]},tooltip:{},radar:{radius:"60%",splitNumber:8,indicator:[{name:"2017"},{name:"2017"},{name:"2018"},{name:"2019"},{name:"2020"},{name:"2021"}]},series:[{type:"radar",symbolSize:0,areaStyle:{shadowBlur:0,shadowColor:"rgba(0,0,0,.2)",shadowOffsetX:0,shadowOffsetY:10,opacity:1},data:[{value:[90,50,86,40,50,20],name:"Visits",itemStyle:{color:"#b6a2de"}},{value:[70,75,70,76,20,85],name:"Sales",itemStyle:{color:"#67e0e3"}}]}]})},{immediate:!0}),(p,g)=>(l(),n(f(c),{title:"\u9500\u552E\u7EDF\u8BA1",loading:e.loading},{default:d(()=>[m("div",{ref_key:"chartRef",ref:t,style:u({width:e.width,height:e.height})},null,4)]),_:1},8,["loading"]))}});export{x as default};