import{d as R,k as b,r as w,Q as S,f as s,c as t,t as n,h as r,P as g,M as $,e as x,o as l,R as p,z as o,J as V,I as G,B as k,K as _,F as h,m as T,_ as B}from"./index-nw_d5X0a.js";import{s as D,d as E}from"./traffic-mk_vra8f.js";import{S as O,a as A,s as f}from"./SearchUtil-xjq24Jx3.js";import"./request-hwqjhTR7.js";const M={class:"__container_resources_application_index"},P=["onClick"],F=R({__name:"index",setup(Y){b(e=>({"483ed2f8":r(g)}));let N=[{title:"ruleName",key:"ruleName",dataIndex:"ruleName",sorter:(e,a)=>f(e.appName,a.appName),width:140},{title:"ruleGranularity",key:"ruleGranularity",dataIndex:"ruleGranularity",render:(e,a)=>a.isService?"服务":"应用",width:100,sorter:(e,a)=>f(e.instanceNum,a.instanceNum)},{title:"createTime",key:"createTime",dataIndex:"createTime",width:120,sorter:(e,a)=>f(e.instanceNum,a.instanceNum)},{title:"enable",key:"enable",dataIndex:"enable",render:(e,a)=>a.enable?"是":"否",width:120,sorter:(e,a)=>f(e.instanceNum,a.instanceNum)},{title:"operation",key:"operation",dataIndex:"operation",width:200}];const c=w(new O([{label:"serviceGovernance",param:"serviceGovernance",placeholder:"typeRoutingRules",style:{width:"200px"}}],D,N)),C=async e=>{(await E(e)).code===200&&await c.onSearch()};S(()=>{c.onSearch()});const v=e=>{C(e)};return $(T.SEARCH_DOMAIN,c),(e,a)=>{const d=x("a-button"),I=x("a-popconfirm");return l(),s("div",M,[t(A,{"search-domain":c},{customOperation:n(()=>[t(d,{type:"primary",onClick:a[0]||(a[0]=u=>r(p).push("/traffic/addRoutingRule"))},{default:n(()=>[o("新增条件路由规则 ")]),_:1})]),bodyCell:n(({text:u,column:i,record:m})=>[i.dataIndex==="ruleName"?(l(),s("span",{key:0,class:"rule-link",onClick:y=>r(p).push(`formview/${m[i.key]}`)},[V("b",null,[t(r(G),{style:{"margin-bottom":"-2px"},icon:"material-symbols:attach-file-rounded"}),o(" "+k(u),1)])],8,P)):_("",!0),i.dataIndex==="ruleGranularity"?(l(),s(h,{key:1},[o(k(u?"服务":"应用"),1)],64)):_("",!0),i.dataIndex==="enable"?(l(),s(h,{key:2},[o(k(u?"启用":"禁用"),1)],64)):_("",!0),i.dataIndex==="operation"?(l(),s(h,{key:3},[t(d,{type:"link",onClick:y=>r(p).push(`formview/${m.ruleName}`)},{default:n(()=>[o(" 查看 ")]),_:2},1032,["onClick"]),t(d,{type:"link",onClick:y=>r(p).push(`/traffic/updateRoutingRule/updateByFormView/${m.ruleName}`)},{default:n(()=>[o(" 修改 ")]),_:2},1032,["onClick"]),t(I,{title:"确认删除该条件路由规则？","ok-text":"Yes","cancel-text":"No",onConfirm:y=>v(m.ruleName)},{default:n(()=>[t(d,{type:"link"},{default:n(()=>[o(" 删除")]),_:1})]),_:2},1032,["onConfirm"])],64)):_("",!0)]),_:1},8,["search-domain"])])}}}),L=B(F,[["__scopeId","data-v-5f67b727"]]);export{L as default};
