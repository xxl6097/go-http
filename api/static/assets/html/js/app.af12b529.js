(function(){"use strict";var t={1471:function(t,o,e){var s=e(9242),n=e(3396);function i(t,o,e,s,i,l){const a=(0,n.up)("LogView");return(0,n.wg)(),(0,n.iD)("div",null,[(0,n.Wm)(a)])}var l=e(7139);const a=t=>((0,n.dD)("data-v-100c2d4c"),t=t(),(0,n.Cn)(),t),h={class:"homewrap"},c={style:{"text-align":"left",border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)",padding:"5px","margin-top":"1px"}},r={style:{display:"flex",width:"100%","margin-top":"5px"}},g={style:{"text-align":"left",border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04)",padding:"5px","margin-top":"15px"}},d=a((()=>(0,n._)("div",{class:"border-title"},[(0,n._)("span",null,"文件结构")],-1))),p=["onClick"],u={style:{width:"100%"}},w={style:{border:"solid 1px #d9dede","box-shadow":"0 2px 4px rgba(0, 0, 0, 0.12),\n                0 0 6px rgba(0, 0, 0, 0.04)",padding:"5px",width:"100%",height:"90%",overflow:"auto","word-break":"break-all"}},f={ref:"txtContent",id:"txtContent",style:{"text-align":"left",border:"1px solid #646060",height:"100%",overflow:"auto","word-break":"break-all"}};function b(t,o,e,s,i,a){const b=(0,n.up)("el-input"),v=(0,n.up)("el-form-item"),m=(0,n.up)("el-button"),x=(0,n.up)("el-form"),y=(0,n.up)("el-scrollbar"),k=(0,n.up)("el-aside"),D=(0,n.up)("el-container"),_=(0,n.Q2)("loading");return(0,n.wg)(),(0,n.iD)("div",h,[(0,n.Wm)(D,{class:"home-container"},{default:(0,n.w5)((()=>[(0,n.wy)(((0,n.wg)(),(0,n.j4)(k,{width:"20%"},{default:(0,n.w5)((()=>[(0,n._)("div",c,[(0,n.Wm)(x,null,{default:(0,n.w5)((()=>[(0,n.Wm)(v,{label:"ws地址："},{default:(0,n.w5)((()=>[(0,n.Wm)(b,{size:"mini",modelValue:t.wshost,"onUpdate:modelValue":o[0]||(o[0]=o=>t.wshost=o),placeholder:"请输入ws地址"},null,8,["modelValue"])])),_:1}),(0,n._)("div",r,[(0,n.Wm)(m,{type:t.datas.btncolor,style:{"font-size":"10px",width:"100%"},onClick:o[1]||(o[1]=o=>t.onStart())},{default:(0,n.w5)((()=>[(0,n.Uk)((0,l.zw)(t.datas.btntext),1)])),_:1},8,["type"]),(0,n.Wm)(m,{size:"mini",style:{"font-size":"10px",width:"100%"},onClick:o[2]||(o[2]=o=>t.onRefreshLog())},{default:(0,n.w5)((()=>[(0,n.Uk)("刷新日志")])),_:1}),(0,n.Wm)(m,{size:"mini",style:{"font-size":"10px",width:"100%"},onClick:o[3]||(o[3]=o=>t.onClearLog())},{default:(0,n.w5)((()=>[(0,n.Uk)("清空日志")])),_:1})])])),_:1})]),(0,n._)("div",g,[d,(0,n.Wm)(y,{height:"600px"},{default:(0,n.w5)((()=>[(0,n._)("a",{href:"#",onClick:o[4]||(o[4]=o=>t.onFileClick(".."))},".."),((0,n.wg)(!0),(0,n.iD)(n.HY,null,(0,n.Ko)(t.files,(o=>((0,n.wg)(),(0,n.iD)("div",{key:o.id},[(0,n._)("a",{href:"#",onClick:e=>t.onFileClick(o.label)},(0,l.zw)(o.label),9,p)])))),128))])),_:1})])])),_:1})),[[_,t.loading]]),(0,n.Wm)(D,null,{default:(0,n.w5)((()=>[(0,n._)("div",u,[(0,n._)("div",w,[(0,n._)("div",f,null,512)])])])),_:1})])),_:1})])}var v=e(7327),m=e(6520),x=e(6e3),y=e(7178),k=function(t,o,e,s){var n,i=arguments.length,l=i<3?o:null===s?s=Object.getOwnPropertyDescriptor(o,e):s;if("object"===typeof Reflect&&"function"===typeof Reflect.decorate)l=Reflect.decorate(t,o,e,s);else for(var a=t.length-1;a>=0;a--)(n=t[a])&&(l=(i<3?n(l):i>3?n(o,e,l):n(o,e))||l);return i>3&&l&&Object.defineProperty(o,e,l),l};const D=(t,o)=>`${t}-${o}`,_=(t,o,e,s=1,n="node")=>{let i=0;return Array.from({length:e}).fill(s).map((()=>{const e=s===t?0:Math.round(Math.random()*o),l=D(n,++i);return{id:l,label:l,children:e?_(t,o,e,s+1,l):void 0}}))};let C=class extends m.w3{constructor(...t){super(...t),(0,v.Z)(this,"msg",void 0),(0,v.Z)(this,"websock",void 0),(0,v.Z)(this,"props",{value:"id",label:"label",children:"children"}),(0,v.Z)(this,"loading",!1),(0,v.Z)(this,"connStatus",!1),(0,v.Z)(this,"options",[{value:"wss://cyjy-iot.chengyang.gov.cn/clink/gtbx/logws",label:"城阳环境"},{value:"ws://iot.clife.net:31307/echo",label:"clife生产环境"}]),(0,v.Z)(this,"treeData",_(2,2,3)),(0,v.Z)(this,"datas",{btntext:"打开",btncolor:"primary",websock:null,logDiv:document.getElementById("txtContent")}),(0,v.Z)(this,"wshost",""),(0,v.Z)(this,"apihost",window.location.origin),(0,v.Z)(this,"files",[]),(0,v.Z)(this,"logpath","/")}onStart(){if(""===this.wshost)return x.T.alert("请填写ws地址哦～～","警告",{confirmButtonText:"OK",callback:t=>{(0,y.z8)({type:"info",message:`action: ${t}`})}}),console.log("地址是空的哦");console.log("打开"),this.loading=!0,this.connStatus?this.websock.close():this.initWebSocket()}initWebSocket(){if(console.log("initWebSocket"),"undefined"===typeof WebSocket)return console.log("您的浏览器不支持websocket");console.log("host",this.wshost);try{this.websock=new WebSocket(this.wshost),this.websock.onopen=t=>{this.datas.btncolor="danger",this.loading=!1,this.connStatus=!0,console.log(t.currentTarget.url),console.log("websocket connect sucessully..",t),this.showLog("连接成功 "+t.currentTarget.url),this.datas.btntext="关闭"},this.websock.onmessage=t=>{console.log("onmessage",t),this.showLog(t.data)},this.websock.onclose=t=>{this.datas.btncolor="primary",this.loading=!1,console.log("onclose received a message",t),this.showLog("连接关闭:"+JSON.stringify(t)),this.connStatus=!1,this.datas.btntext="打开"},this.websock.onerror=t=>{this.datas.btncolor="primary",this.loading=!1,console.log("onerror received a message",t),this.showLog("连接错误:"+JSON.stringify(t)),this.connStatus=!1,this.datas.btntext="打开"}}catch(t){console.log("catch received a message",t)}}showLog(t){var o=document.createElement("div");o.append(t),null!=this.datas.logDiv&&(this.datas.logDiv.append(o),this.datas.logDiv.scrollTop=this.datas.logDiv.scrollHeight,this.datas.logDiv.scrollIntoView())}onClearLog(){null!=this.datas.logDiv&&(this.datas.logDiv.innerText="")}onRefreshLog(){this.logpath="/",this.fetchData(this.logpath)}onFileClick(t){if(console.log("onFileClick",t,this.logpath),".."===t){if("/"===this.logpath)return;let t=this.logpath.split("/");console.log("logpath",t);let o="";t.forEach(((t,e,s)=>{e<s.length-2&&(console.log("forEach",t,e,s.length),o=o.concat(t,"/"))})),this.logpath=o}else{if(console.log("logpath e",t),console.log("logpath ->",this.apihost+"fserver"+this.logpath+t),!t.endsWith("/"))return console.log("logpath e2",this.apihost+"/fserver"+this.logpath+t),void window.open(this.apihost+"/fserver/"+this.logpath+t,"_blank");this.logpath=this.logpath.concat(t),console.log("logpath e1",this.logpath)}this.fetchData(this.logpath)}get allname(){return"computed "+this.msg}fetchData(t){console.log("fetchData",t);let o="";o=o.concat(this.apihost,window.location.pathname,"api/files?path=",t),this.showLog(o),console.log("apipath",o),fetch(o,{credentials:"include"}).then((t=>{let o=t.headers.get("File-Type");return console.log("1-fetchData",t,t.statusText,o),"text"==o?(console.log("2-fetchData",t),t.text()):t.json()})).then((t=>{console.log("4-fetchData",t),this.files=t})).catch((t=>{(0,y.z8)({showClose:!0,message:"Get status failed!"+t,type:"warning"}),this.showLog(t),console.log("3-fetchData",t)}))}fetchDirs(){fetch("http://localhost:8080/fserver/",{credentials:"include"}).then((t=>(this.showLog(JSON.stringify(t)),t.json()))).then((t=>{this.treeData=t,this.showLog(JSON.stringify(t))})).catch((t=>{(0,y.z8)({showClose:!0,message:"Get status failed!"+t,type:"warning"}),this.showLog(t)}))}testFetch(){fetch("http://localhost:8081/v1/api/device/all").then((t=>(this.showLog(JSON.stringify(t)),t.json()))).then((t=>{this.showLog(JSON.stringify(t))})).catch((t=>{(0,y.z8)({showClose:!0,message:"Get status failed!"+t,type:"warning"}),this.showLog(t)}))}created(){console.log("created")}mounted(){this.datas.logDiv=document.getElementById("txtContent"),console.log("mounted",window.location),console.log("host",window.location.host),console.log("origin",window.location.origin),console.log("pathname",window.location.pathname),console.log("protocol",window.location.protocol);let t=window.location.pathname,o=t.split("/");console.log("list",o);let e="";o.forEach(((t,o,s)=>{o>0&&o<s.length-2&&(console.log("forEach",t,o,s.length),e=e.concat(t,"/"))})),console.log("api",e);let s=window.location.protocol;s.startsWith("https")?this.wshost=this.wshost.concat("wss://",window.location.host,e):this.wshost=this.wshost.concat("ws://",window.location.host,e),this.wshost=this.wshost.concat(window.location.pathname),this.wshost=this.wshost.concat("echo"),console.log("wshost",this.wshost),this.apihost=this.apihost.concat(e),console.log("apihost",this.apihost),this.showLog(JSON.stringify(this.treeData)),this.fetchData(this.logpath)}};C=k([(0,m.Ei)({props:{msg:String}})],C);var O=C,L=e(89);const S=(0,L.Z)(O,[["render",b],["__scopeId","data-v-100c2d4c"]]);var W=S,Z=(0,n.aZ)({name:"HomeView",components:{LogView:W}});const j=(0,L.Z)(Z,[["render",i]]);var z=j,T=e(5280);e(4415);(0,s.ri)(z).use(T.Z).mount("#app")}},o={};function e(s){var n=o[s];if(void 0!==n)return n.exports;var i=o[s]={exports:{}};return t[s].call(i.exports,i,i.exports,e),i.exports}e.m=t,function(){var t=[];e.O=function(o,s,n,i){if(!s){var l=1/0;for(r=0;r<t.length;r++){s=t[r][0],n=t[r][1],i=t[r][2];for(var a=!0,h=0;h<s.length;h++)(!1&i||l>=i)&&Object.keys(e.O).every((function(t){return e.O[t](s[h])}))?s.splice(h--,1):(a=!1,i<l&&(l=i));if(a){t.splice(r--,1);var c=n();void 0!==c&&(o=c)}}return o}i=i||0;for(var r=t.length;r>0&&t[r-1][2]>i;r--)t[r]=t[r-1];t[r]=[s,n,i]}}(),function(){e.n=function(t){var o=t&&t.__esModule?function(){return t["default"]}:function(){return t};return e.d(o,{a:o}),o}}(),function(){e.d=function(t,o){for(var s in o)e.o(o,s)&&!e.o(t,s)&&Object.defineProperty(t,s,{enumerable:!0,get:o[s]})}}(),function(){e.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){e.o=function(t,o){return Object.prototype.hasOwnProperty.call(t,o)}}(),function(){var t={143:0};e.O.j=function(o){return 0===t[o]};var o=function(o,s){var n,i,l=s[0],a=s[1],h=s[2],c=0;if(l.some((function(o){return 0!==t[o]}))){for(n in a)e.o(a,n)&&(e.m[n]=a[n]);if(h)var r=h(e)}for(o&&o(s);c<l.length;c++)i=l[c],e.o(t,i)&&t[i]&&t[i][0](),t[i]=0;return e.O(r)},s=self["webpackChunkvue_weblog"]=self["webpackChunkvue_weblog"]||[];s.forEach(o.bind(null,0)),s.push=o.bind(null,s.push.bind(s))}();var s=e.O(void 0,[998],(function(){return e(1471)}));s=e.O(s)})();