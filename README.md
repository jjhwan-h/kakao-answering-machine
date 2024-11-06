<img width="568" alt="image" src="https://github.com/user-attachments/assets/36cf45e2-f35f-4e57-b751-7c73a125267e">

# 개요
유튜브를 보다가 빵형의 개발도상국님의 카카오톡 매크로(https://www.youtube.com/watch?v=tFVk-o2JA3o)를 보게되었다.<br>
재밌을것 같아서 python대신 golang으로 만들어 보았다.

# 기능
v1: 날씨를 알려주는 기능 구현. (모든 것은 자동으로 이루어진다.)
1. 1분마다 카카오톡 채팅창을 스캔
2. 새로운 채팅 알림이 뜰경우 빨간색 원이 나타난다.
   <img width="64" alt="image" src="https://github.com/user-attachments/assets/d402f442-79d5-4093-8be6-79052431b867">
3. 이를 인식하여 빨간색원의 좌표를 더블클릭하여 채팅방 접속
4. 채팅방하단의 내용을 캡쳐하여 OCR서버로 전송
5. 날씨 [지역]으로된 문장이 있을 경우 [지역]만 리턴
6. OpenWeather(https://openweathermap.org/current)의 current weather API를 이용하여 [지역]의 날씨를 받아온다.
7. 받아온 날씨 데이터를 다듬어서 채팅방에 제공

# 실행환경
- macOS: Sonoma 14.2.1
- golang: 1.22.2
- python(OCR서버): 3.10
  
# 실행
0. setup
  - .env
  ```
  CURRENTAPI= [ openweathermap api url]
  KEY= [ openweathermap api key]
  LANG="kr"
  OCRAPI= [ ocr 서버url]
  ```
  - 채팅대기방 채팅창 위치
    - 새로운 채팅 스캔 및 채팅 내용 캡쳐 시 채팅방의 위치가 중요
    - cmd/serve.go의 Start에서 상황에 따라 수정
     ```
        func Start() {
      	messageBox := &internal.DIR{X1: 490, Y1: 900, X2: 570, Y2: 955}
      	speechBox := &internal.SpeechBox{X: 0, Y: 760, Width: 560, Height: 72}
     ```
    - messageBox는 채팅대기방의 box<br>
    - speechBox는 채팅방에서 캡쳐할 위치. 즉, 새로운 메시지가 있을것으로 예상되는 부분<br>
    
1. Open weather의 API 키가 필요
https://openweathermap.org/current에 접속하여 회원가입 후 Current weather data API(무료)신청 후 API키 발급받는다.<br>

2. 정확한 지역명을 위해 매칭 테이블을 만든다.
   외국에서 제공하는 API라 영어로 된 지역명이 정확히지 않음 (예를들어, 진천 => Chinch'ŏn)
   ```go
     package internal

      var (
    	CITYTABLE = map[string]string{
    		"진천": "Chinch'ŏn",
    		"서울": "Seoul",
    		"울산": "Ulsan",
    		"대전": "Daejeon",
    	}
    	REVERSCITYTABLE = map[string]string{
    		"Chinch'ŏn": "진천",
    		"Seoul":     "서울",
    		"Ulsan":     "울산",
    		"Daejeon":   "대전",
    	}
    )

    다른 도시를 더 추가할 경우 https://bulk.openweathermap.org/sample/city.list.json.gz에서 확인 후 **city_table.go**파일에 추가

3. 이전에 만들었던 OCR서버를 이용
   https://github.com/jjhwan-h/2023-CBNU-OpenSourceProject
    - build
    ```
      docker build -t ocr .
    ```
    - run
    ```
      docker run -p 3000:5000 ocr
    ```
4. bot 실행
   ```
      go run main.go serve
   ```

# 에러
- 혹시라도 robotgo(https://github.com/go-vgo/robotgo)때문에 에러가 발생한다면 xcode설치

# 실행영상


# v2
- chatgpt API를 이용해서 날씨뿐아니라 일상대화도 자동으로 할 수 있게 바꿀예정

