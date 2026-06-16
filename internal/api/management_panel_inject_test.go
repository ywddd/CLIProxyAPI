package api

import (
	"bytes"
	"testing"
)

func TestPatchManagementAuthFilesFilters(t *testing.T) {
	input := []byte("aT=(e,t)=>{let n=z_(iT(e,t));return n?n===`pro`?50:Xw.has(n)&&n!==`pro`?40:n===`team`?30:n===`plus`?20:n===`free`?10:0:null},oT=" +
		"[c,l]=(0,y.useState)(`all`),[u,d]=(0,y.useState)(!1),[f,p]=(0,y.useState)(!1),[m,h]=(0,y.useState)(!1),[g,_]=(0,y.useState)(!1)," +
		"typeof t.healthyOnly==`boolean`&&h(t.healthyOnly),typeof e!=`boolean`&&typeof t.compactMode==`boolean`&&_(t.compactMode)," +
		"healthyOnly:m,compactMode:g,search:v,page:x,pageSize:at,regularPageSize:w.regular,compactPageSize:w.compact,sortMode:A,viewMode:O}" +
		"},[g,f,c,m,x,at,w,u,v,A,P,O])" +
		"pt=(0,y.useMemo)(()=>ne.filter(e=>!(u&&!Xx(e)||f&&e.disabled!==!0||m&&!Zx(e))),[f,ne,m,u])" +
		"{value:`plan-desc`,label:e(`auth_files.sort_plan_desc`)},{value:`plan-asc`,label:e(`auth_files.sort_plan_asc`)}" +
		"A===`priority-asc`||A===`priority-desc`?e.sort((e,t)=>tT(e,t,A===`priority-desc`?`desc`:`asc`)):(A===`plan-asc`||A===`plan-desc`)" +
		"(0,H.jsx)(`div`,{className:J.filterToggleCard,children:(0,H.jsx)(Dy,{checked:g,onChange:e=>_(e),ariaLabel:e(`auth_files.compact_mode_label`)")

	out := patchManagementAuthFilesFilters(input)

	for _, want := range [][]byte{
		[]byte("cpaIsFreeAuth"),
		[]byte("cpaFreeOnly"),
		[]byte("freeOnly:cpaFreeOnly"),
		[]byte("created-desc"),
		[]byte("cpaAuthTime"),
		[]byte("\\u663e\\u793afree\\u8d26\\u53f7"),
		[]byte("\\u663e\\u793aplus\\u8d26\\u53f7"),
	} {
		if !bytes.Contains(out, want) {
			t.Fatalf("patched output missing %q", want)
		}
	}
	for _, forbidden := range [][]byte{
		[]byte("e.name??"),
		[]byte("iT(e,t)??"),
		[]byte("cpaIsFreeAuth(e,i[e.name])"),
	} {
		if bytes.Contains(out, forbidden) {
			t.Fatalf("patched output should not use name-derived plan text %q", forbidden)
		}
	}
}
