# CPA 鑷敤鐗?

杩欐槸鍩轰簬 [router-for-me/CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) 鐨勮嚜鐢ㄦ瀯寤猴紝閲嶇偣鏈嶅姟 Codex/Responses 绋冲畾鎬с€佸璐﹀彿杩愯銆丯AS/Docker 閮ㄧ讲鍜屾棩甯?CPA 绠＄悊銆?

褰撳墠鍚屾鍩虹嚎锛氫笂娓?`v7.2.43`銆傝嚜鐢ㄧ増鏈缓璁爣璁颁负锛?

```text
v7.2.43-selfuse.20260628
```

## 鏈瀯寤轰繚鐣欑殑鏀瑰姩

### 1. Codex 涓婁笅鏂囪繃闀跨洿鎺ヤ氦鍥炲鎴风

褰?Codex 涓婃父浠?`context_too_large` / `context_length_exceeded` 缁撴潫娴佸紡鍝嶅簲鏃讹紝鏈瀯寤轰笉鍦?CPA 涓棿灞傝嚜琛屽帇缂╁巻鍙层€佺敓鎴?`history.txt` 鎴栫Щ闄?reasoning 鍚庣户缁噸璇曘€?

杩欐牱鍙互閬垮厤 CPA 鎶婂巻鍙蹭細璇濇敼鍐欐垚鏂扮殑璇锋眰鍚庡啀娆″杺缁欐ā鍨嬶紝闄嶄綆闀夸細璇濋噷閲嶅璇诲伐浣滃尯銆侀噸澶嶇‘璁ょ姸鎬併€侀噸澶嶈鍒掔殑椋庨櫓銆?

### 2. 鍔犲瘑 reasoning 涓婁笅鏂囬檷绾ч噸璇?

閮ㄥ垎 Codex/Responses 璇锋眰浼氭惡甯?`input[*].encrypted_content`銆傚綋涓婃父鏄庣‘鎷掔粷杩欐鍔犲瘑 reasoning 涓婁笅鏂囨椂锛屾湰鏋勫缓浼氱Щ闄ゆ棤鏁堢殑鍔犲瘑 reasoning 涓婁笅鏂囷紝骞堕噸璇曚竴娆°€?

鍚屾椂锛屽綋涓婃父杩斿洖 `Item with id 'rs_...' not found` 涓旀彁绀?`store=false` 鏃讹紝涔熶細绉婚櫎 stale reasoning item 骞堕噸璇曚竴娆°€?

### 3. Codex 鍝嶅簲澶磋秴鏃?

Codex 涓婃父璇锋眰鏈夋椂浼氬湪杩斿洖鍝嶅簲澶村墠鍗′綇銆傛湰鏋勫缓澧炲姞鍙綔鐢ㄤ簬鍝嶅簲澶撮樁娈电殑瓒呮椂锛?

```yaml
codex-response-header-timeout-seconds: 180
```

鍝嶅簲澶村埌杈惧悗鐨勬祦寮忔鏂囦笉鍙楄瓒呮椂闄愬埗銆傝缃负璐熸暟鍙叧闂細

```yaml
codex-response-header-timeout-seconds: -1
```

涔熸敮鎸佺幆澧冨彉閲忚鐩栵細

```bash
CPA_CODEX_RESPONSE_HEADER_TIMEOUT_SECONDS=180
```

### 4. OpenAI-compatible JSON 棰勬

Kimi K2.7 Code 绛夎蛋 `openai-compatibility` 鐨勬ā鍨嬪湪璇锋眰浣撳寘鍚湭杞箟鍙嶆枩鏉犳椂锛屼笂娓稿彲鑳借繑鍥?Cloudflare 渚х殑 `invalid escaped character in string`銆?

鏈瀯寤轰細鍦ㄥ叆鍙ｈ矾鐢卞墠鍜屽彂寰€ OpenAI-compatible 涓婃父鍓嶅 JSON 鍋氬吋瀹瑰鐞嗭細

- 瀵?`C:\Users\...` 杩欑被甯歌鏈浆涔?Windows 璺緞锛岃嚜鍔ㄤ慨澶嶅瓧绗︿覆閲岀殑闈炴硶鍙嶆枩鏉犺浆涔夊悗缁х画璇锋眰銆?
- `/v1/chat/completions` 鍜?`/v1/completions` 浼氬厛淇/鏍￠獙璇锋眰浣擄紝鍐嶈鍙?`model` 鍋?provider 璺敱銆?
- 瀵圭己寮曞彿銆佺粨鏋勬崯鍧忕瓑涓嶅彲鎭㈠鐨勯潪娉?JSON锛屼粛鐒跺湪 CPA 鏈湴杩斿洖 `400`銆?

### 5. 绠＄悊 UI 澧炲己

绠＄悊椤典繚鐣?selfuse 鐨勮繍缁村寮猴細

- 鍙鍖栭厤缃?`codex-response-header-timeout-seconds`銆?
- auth 鏂囦欢鍗曠嫭娴嬭瘯妯″瀷銆?
- 褰撳墠椤垫壒閲忔祴璇?auth 鏂囦欢銆?
- 姣忎釜璐﹀彿鏄剧ず娴嬭瘯缁撴灉鍜屽欢杩熴€?

## 涓婃父鍚屾鎽樿

鏈疆浠?`v7.2.16` 鍚堝苟鍒?`v7.2.43`锛岄噸鐐瑰寘鎷細

- 绠＄悊鏃ュ織 API 澧炲姞 cursor/tail/杞浆缁鑳藉姏銆?
- 鎻掍欢鍒犻櫎銆侀厤缃慨鏀广€佺敓鍛藉懆鏈熷拰 stream callback 鐨勫紓姝?reload/race 淇銆?
- 鏂板鎻掍欢 ModelRouter锛屽彲鍦ㄩ壌鏉冨墠鍋氭ā鍨嬭矾鐢便€?
- Claude/Anthropic 鍏煎澧炲己锛屽寘鎷?web search tool domain 娓呮礂銆乼ool_result 瑙勮寖鍖栥€丆odex web_search_call 娴佸紡杞崲淇銆乶amespace/function call 鏄犲皠澧炲己銆?
- 瑙嗛杈撳叆澧炲己锛屽鍔?`video_url` 鎻愬彇鍜屾牎楠屻€?
- 鎻掍欢榛樿 `Enabled` 琛屼负鏀逛负 `false`锛屽凡鏈夋彃浠堕厤缃渶瑕佹樉寮忓惎鐢ㄣ€?

涓婃父宸茬粡瑕嗙洊鐨勯€氱敤淇灏介噺浣跨敤瀹樻柟瀹炵幇锛涗笂娓稿皻鏈鐩栫殑 selfuse 杩愯琛ヤ竵缁х画淇濈暀銆?

## 鎺ㄨ崘閰嶇疆

```yaml
request-retry: 3
max-retry-credentials: 3
max-retry-interval: 30

routing:
  session-affinity: true

nonstream-keepalive-interval: 15
codex-response-header-timeout-seconds: 180

streaming:
  keepalive-seconds: 15
  bootstrap-retries: 1
```

## Docker Compose 浣跨敤

鏋勫缓骞跺惎鍔細

```bash
docker compose up -d --build
```

绠＄悊椤靛拰 API 绔彛鍙栧喅浜庝綘鐨?compose 鏂囦欢銆傚弬鑰冮儴缃蹭腑锛?

```text
CPA API:    http://<host>:8317
CPA Plus:   http://<host>:18317/management.html
```

## 鐗堟湰瑙勫垯

鏈粨搴撶殑鑷敤鍙戝竷鐗堟湰鍥哄畾浣跨敤 `selfuse` 鍚庣紑锛屼緥濡傦細

```text
v7.2.43-selfuse.20260628
```

NAS 鏈湴 Docker 闀滃儚寤鸿浣跨敤绋冲畾鏍囩锛?

```text
cli-proxy-api:v7.2.43-selfuse.20260628
```

## 瀹夊叏璇存槑

涓嶈鎻愪氦鐪熷疄 auth 鏂囦欢銆乺efresh token銆乤ccess token銆乮d token銆乵anagement key 鎴?API key銆傛帹鑽愪綔涓鸿繍琛屾€佹枃浠朵繚鐣欏湪浠撳簱澶栨垨 `.gitignore` 涓細

```text
auth-dir/
auths/
logs/
*.sqlite
*.db
config.yaml
```

鍏紑 fork 鎴栧彂甯冨墠锛屽缓璁壂鎻忔晱鎰熶俊鎭細

```bash
rg -n "github_pat_|refresh_token|access_token|id_token|sk-[A-Za-z0-9]|secret-key:" .
```
