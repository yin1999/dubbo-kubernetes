import{d as l,l as m,m as d,n as _,u as f,Q as u,e as p,o as r,f as s,c as h,t as g,J as y,h as n,K as v,_ as I}from"./index-nw_d5X0a.js";const w={class:"__container_tabDemo3"},b={class:"__container_iframe_container"},q=["src"],S=l({__name:"GrafanaPage",setup(k){const a=m(d.GRAFANA);_(""),f(),u(async()=>{var t;let e=await a.api({});a.url=`${window.location.origin}/grafana/d/${(t=e.data)==null?void 0:t.baseURL.split("/d/")[1].split("?")[0]}?var-${a.type}=${a.name}&kiosk=tv`,a.showIframe=!0});function o(e){try{e()}catch(t){console.log(t)}}function c(){console.log("The iframe has been loaded."),setTimeout(()=>{try{let e=document.querySelector("#grafanaIframe").contentDocument;o(()=>{e.querySelector("header").remove()}),o(()=>{e.querySelector("[data-testid*='controls']").remove()}),setTimeout(()=>{o(()=>{e.querySelector("[data-testid*='navigation mega-menu']").remove()}),o(()=>{for(let t of e.querySelectorAll("[data-testid*='Panel menu']"))t.remove()})},2e3)}catch{}a.showIframe=!0},1e3)}return(e,t)=>{const i=p("a-spin");return r(),s("div",w,[h(i,{class:"spin",spinning:!n(a).showIframe},{default:g(()=>[y("div",b,[n(a).showIframe?(r(),s("iframe",{key:0,onload:c,id:"grafanaIframe",style:{"padding-top":"60px"},src:n(a).url,frameborder:"0"},null,8,q)):v("",!0)])]),_:1},8,["spinning"])])}}}),A=I(S,[["__scopeId","data-v-c6fbb32b"]]);export{A as G};
