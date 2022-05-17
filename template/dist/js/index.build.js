/*! For license information please see index.build.js.LICENSE.txt */
"use strict";(globalThis.webpackChunkcustom_cli=globalThis.webpackChunkcustom_cli||[]).push([[826],{7381:(e,t,r)=>{var n=r(3010),o=r(9529),s=r(5342),i=r(6739),l=r(5170),c=r(1433),a=r(6562),u=r(4032),d=r(1185);const p=(0,l.createRef)(),h=()=>{const[e,t]=(0,l.useState)(!1),r=(0,l.useRef)(null),[o,s]=(0,l.useState)(""),[i,h]=(0,l.useState)("info"),[x,m]=(0,l.useState)({vertical:"bottom",horizontal:"right"}),[g,v]=(0,l.useState)(void 0),[j,b]=(0,l.useState)(3e3);(0,l.useEffect)((()=>(e&&(r.current=setTimeout((()=>{t(!1)}),j)),()=>{clearTimeout(r.current),r.current=null})),[e,i,x]);const f={left:e=>(0,n.jsx)(d.Z,{...e,direction:"left"}),up:e=>(0,n.jsx)(d.Z,{...e,direction:"up"}),right:e=>(0,n.jsx)(d.Z,{...e,direction:"right"}),down:e=>(0,n.jsx)(d.Z,{...e,direction:"down"})},y=(0,l.useCallback)((()=>t(!1)),[]),Z=(0,l.useCallback)((e=>{let{message:r,type:n,origin:o,duration:i,direction:l="down"}=e;n&&h(n),o&&m(o),v((()=>f[l])),s(r),i&&b(i),t(!0)}),[]);return(0,l.useImperativeHandle)(p,(()=>({close:y,open:Z})),[y,Z]),(0,n.jsx)(c.Z,{message:o,open:e,anchorOrigin:x,TransitionComponent:g,children:(0,n.jsxs)(a.Z,{onClose:y,severity:i,children:[(0,n.jsx)(u.Z,{children:i}),"This is a ",i," alert —",(0,n.jsx)("strong",{children:o})]})})};var x=r(365),m=r.n(x);function g(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}const v=new class{constructor(e){var t,r,n,o;g(this,"instance",void 0),g(this,"interceptorsObj",void 0),this.instance=m().create(e),this.interceptorsObj=e.intercepotrs,this.instance.interceptors.request.use((e=>e),(e=>e)),this.instance.interceptors.request.use(null===(t=this.interceptorsObj)||void 0===t?void 0:t.requestInterceptors,null===(r=this.interceptorsObj)||void 0===r?void 0:r.requestInterceptorsCatch),this.instance.interceptors.response.use(null===(n=this.interceptorsObj)||void 0===n?void 0:n.responseInterceptors,null===(o=this.interceptorsObj)||void 0===o?void 0:o.responseInterceptorsCatch),this.instance.interceptors.response.use((e=>e.data),(e=>e))}request(e){return this.instance.request(e)}get(e){return this.request({method:"get",...e})}post(e){return this.request({method:"post",...e})}put(e){return this.request({method:"put",...e})}delete(e){return this.request({method:"delete",...e})}}({baseURL:"http://localhost:9000",timeout:3e3,intercepotrs:{requestInterceptors:e=>e,requestInterceptorsCatch:e=>e,responseInterceptors(e){var t;let r=e.data.code;return null===(t=p.current)||void 0===t||t.open({message:e.data.msg,type:200!==r?"error":"success",origin:{horizontal:"center",vertical:"top"}}),e.data},responseInterceptorsCatch(e){var t;return null===(t=p.current)||void 0===t||t.open({message:e.message,type:"error",origin:{horizontal:"center",vertical:"top"}}),e}}});var j=r(3633),b=r(474),f=r(7726),y=r(4832),Z=r(1943),w=r(3561),C=r(1678),S=r(4071),O=r(9347),E=r(138),k=r(1548),z=r(9613);function I(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function _(e){return null==e}const T=e=>{let{onClick:t,state:r,loading:o}=e;return(0,n.jsx)(z.Z,{sx:{display:"flex",alignItems:"center"},children:(0,n.jsxs)(z.Z,{sx:{m:1,position:"relative"},children:[o&&(0,n.jsx)(E.Z,{size:50,sx:{color:"#fff",position:"absolute",top:"50%",left:"50%",marginTop:"-25px",marginLeft:"-25px",zIndex:1}}),(0,n.jsxs)(Z.Z,{color:"primary",sx:{p:"10px",color:"#eee"},"aria-label":"directions",onClick:t,children:["none"===r&&(0,n.jsx)(w.Z,{fontSize:"large",color:"inherit"}),"success"===r&&(0,n.jsx)(O.Z,{fontSize:"large",color:"success"}),"error"===r&&(0,n.jsx)(k.Z,{fontSize:"large",color:"warning"})]})]})})},q=new class{constructor(){var e=this;I(this,"use",(function(t){let r=arguments.length>1&&void 0!==arguments[1]?arguments[1]:500,n=arguments.length>2&&void 0!==arguments[2]&&arguments[2],o=Date.now();return function(){let s=Date.now();for(var i=arguments.length,l=new Array(i),c=0;c<i;c++)l[c]=arguments[c];n&&(t.apply(e,[...l]),n=!1),s-o>=r&&(t.apply(e,[...l]),o=Date.now())}}))}};function R(e){let{value:t,onChange:r,onConfirm:o,onClear:s}=e;const[i,c]=(0,l.useState)(!1),[a,u]=(0,l.useState)("none"),d=(0,l.useMemo)((()=>i),[i]),p=q.use((()=>{c(!0),u("none"),o().then((e=>{c(!1),h(e?"success":"error")})).catch((e=>{c(!1),h("error")}))})),h=e=>{u(e),"error"===e&&setTimeout((()=>{u("none")}),2e3)};return(0,n.jsxs)(b.Z,{component:"form",sx:{p:"2px 4px",boxSizing:"border-box",display:"flex",alignItems:"center",width:"700px",background:{error:{color:"rgba(255, 87, 34,.2)"},none:{color:"rgba(255,255,255,.2)"},success:{color:"rgba(76, 175, 80,.2)"}}[a].color,backdropFilter:"blur(10px)"},children:[(0,n.jsx)(Z.Z,{sx:{p:"10px",color:"#fff"},"aria-label":"url",children:(0,n.jsx)(C.Z,{color:"inherit"})}),(0,n.jsx)(f.ZP,{disabled:d,sx:{ml:1,flex:1,color:"#fff"},placeholder:"Enter Your URL",inputProps:{"aria-label":"enter your url"},value:t,onChange:r}),""!==t?(0,n.jsx)(Z.Z,{color:"primary",sx:{p:"10px",color:"#eee"},"aria-label":"directions",onClick:()=>{s&&s(),c(!1),u("none")},children:(0,n.jsx)(S.Z,{fontSize:"inherit",color:"inherit"})}):"",(0,n.jsx)(y.Z,{sx:{height:28,m:.5},orientation:"vertical"}),(0,n.jsx)(T,{loading:i,state:a,onClick:p})]})}var U=r(6994),D=r(4370),L=r(2718),N=r(13),P=r(2990),A=r(7758);const B=e=>{let{show:t,data:r}=e;const o=e=>{(function(e){return m()({url:e,method:"get",responseType:"blob"})})(e).then((e=>{!function(e,t,r){const n=window.URL.createObjectURL(new Blob([e])),o=document.createElement("a");o.style.display="none",o.href=n;const s=new Date+"-"+t+"."+r;o.setAttribute("download",s),document.body.appendChild(o),o.click(),document.body.removeChild(o)}(e.data,"video","mp4")}))};return(0,n.jsx)(P.Z,{in:t,mountOnEnter:!0,unmountOnExit:!0,children:(0,n.jsx)("div",{className:"video-info",children:(0,n.jsx)(U.Z,{sx:{minWidth:700,boxSizing:"border-box",background:"rgba(0,0,0,0.3)",backdropFilter:"blur(10px)"},children:(0,n.jsxs)(D.Z,{children:[(0,n.jsx)("video",{src:r.path,className:"video-box",controls:!0}),(0,n.jsxs)(z.Z,{sx:{width:"100%",fontSize:"14px",m:"10px 0",backgroundColor:"#5F679C",color:"#fff",borderRadius:"4px",p:2,boxSizing:"border-box",display:"flex",alignItems:"center"},children:[(0,n.jsx)(Z.Z,{color:"primary",sx:{p:"5px 5px ",color:"#eee"},"aria-label":"directions",children:(0,n.jsx)(L.Z,{fontSize:"small",color:"inherit"})}),(0,n.jsx)(y.Z,{sx:{height:28,m:.5},orientation:"vertical"}),(0,n.jsx)("span",{style:{flex:1,margin:"0 10px",wordBreak:"break-all",color:"inherit"},children:null==r?void 0:r.path}),(0,n.jsx)(y.Z,{sx:{height:28,m:.5},orientation:"vertical"}),(0,n.jsx)(Z.Z,{color:"primary",sx:{p:"10px",color:"#eee"},"aria-label":"directions",onClick:()=>{!function(e){let t=document.createElement("input");t.id="copyInput",document.body.appendChild(t),t.setAttribute("style","position:fixed;top:100px;right:100px;opacity:1"),t.value=e;let r=document.getElementById("copyInput");try{var n;r.select(),document.execCommand("Copy"),null===(n=p.current)||void 0===n||n.open({message:"复制成功",type:"success",origin:{horizontal:"center",vertical:"top"}})}catch(e){var o;console.log(e),null===(o=p.current)||void 0===o||o.open({message:"复制失败",type:"error",origin:{horizontal:"center",vertical:"top"}})}t.remove()}(r.path)},children:(0,n.jsx)(N.Z,{fontSize:"small",color:"inherit"})}),(0,n.jsx)(y.Z,{sx:{height:28,m:.5,display:"none"},orientation:"vertical"}),(0,n.jsx)(Z.Z,{color:"primary",sx:{p:"10px",color:"#eee",display:"none"},"aria-label":"directions",onClick:()=>{o(r.path)},children:(0,n.jsx)(A.Z,{fontSize:"small",color:"inherit"})})]}),(0,n.jsx)(a.Z,{severity:"info",sx:{m:"10px 0 0"},children:"如果解析的视频无法播放可尝试在浏览器中直接访问!"})]})})})})},F=(0,j.$j)((e=>({user:e.user.user})))((()=>{const[e,t]=(0,l.useState)(""),[r,o]=(0,l.useState)(null);(0,l.useEffect)((()=>{}));return(0,n.jsx)("div",{className:"home-page",children:(0,n.jsxs)("div",{className:"form-box",children:[(0,n.jsx)(R,{value:e,onChange:e=>{_(r)||o(null),t(e.target.value)},onConfirm:()=>(_(r)||o(null),new Promise(((t,r)=>{var n,s;""===e&&(null===(n=p.current)||void 0===n||n.open({message:"请输入链接",type:"warning",origin:{horizontal:"center",vertical:"top"}}),r(!1));(s={key_words:e},v.get({url:"/parse",params:s})).then((e=>{_(e)||setTimeout((()=>{o(e)}),300),t(!_(e))})).catch((e=>{console.log("error",e),r(!1)}))}))),onClear:()=>{t(""),o(null)}}),(0,n.jsx)(B,{show:!_(r),data:{path:null==r?void 0:r.path}})]})})})),M=()=>(0,n.jsx)("div",{children:"404"}),X=()=>(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(h,{}),(0,n.jsx)(s.VK,{children:(0,n.jsxs)(i.Z5,{children:[(0,n.jsx)(i.AW,{path:"/",element:(0,n.jsx)(F,{})}),(0,n.jsx)(i.AW,{path:"*",element:(0,n.jsx)(M,{})})]})})]});var V=r(8509),W=r(5003);const Y="SET_USER",H="CLEAR_USER",K={user:{id:1e3,uid:100215,name:"lilei",age:20,gender:1},token:null},$=function(){let e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:K,t=arguments.length>1?arguments[1]:void 0,r=Object.assign({},e);switch(t.type){case Y:return void(r.user=t.value);case H:return void(r.user=null);default:return r}},G=(0,V.UY)({user:$}),J=("object"==typeof window&&window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__?window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({}):V.qC)((0,V.md)(W.Z)),Q=(0,V.MT)(G,J),ee=document.getElementById("root");o.createRoot(ee).render((0,n.jsx)(j.zt,{store:Q,children:(0,n.jsx)(X,{})}))}},e=>{e.O(0,[971],(()=>{return t=7381,e(e.s=t);var t}));e.O()}]);