# 📌Kakao-answering-machine

유튜브에서 **빵형의 개발도상국**님의 [카카오톡 매크로 영상](https://www.youtube.com/watch?v=tFVk-o2JA3o)을 보다가 영감을 받아 **Python** 대신 **Go 언어**로 구현해보았습니다.

<div style="display: flex; justify-content: space-around;">
   <img width="45%" height="250px" alt="image" src="https://github.com/user-attachments/assets/41199be7-f888-4f2d-b2a0-d225c429e212">
   <img width="45%" height="250px" alt="image" src="https://github.com/user-attachments/assets/6a40d331-29a1-4c36-9bbe-b8cd6a907574">
</div>

---

## 🚀 기능 (v1: 날씨 정보 제공 및 gpt를 이용한 자동응답 제공)

1. **자동 채팅 감지**: 1초마다 카카오톡 채팅창을 스캔합니다.
2. **새 채팅 알림 인식**: 빨간색 원이 나타나면 이를 인식하여 해당 좌표를 더블클릭하여 채팅방에 접속합니다.  
   <img width="64" alt="red dot" src="https://github.com/user-attachments/assets/d402f442-79d5-4093-8be6-79052431b867">
3. **채팅 내용 캡처 및 분석**: 채팅방 하단의 내용을 캡처하여 OCR 서버로 전송합니다.
4. **지역 추출 및 날씨 조회**: "[지역]뚱" 형식의 문장이 있을 경우, 지역명을 추출합니다. 지역은 CITYTABLE에 등록되어 있어야합니다.
5. **날씨 API 호출**: OpenWeather API를 이용하여 해당 지역의 날씨를 조회합니다.
6. **자동 응답생성**: "[지역]뚱"이 아닌 일반적인 "[문장]뚱"인 경우 문장을 추출합니다.
7. **GPT API 호출**: GPT API를 이용하여 해당 문장에 대한 응답을 생성합니다.
8. **응답 제공**: 받은 날씨 데이터 또는 GPT를 통해 생성한 응답을 채팅방에 제공합니다.

---

## ⚙️ 실행환경

- **운영체제**: macOS Sonoma 14.2.1
- **언어**:
  - Go 1.22.2
  - Python 3.10 (OCR 서버용)

---

## 📝 설치 및 실행

### 1. 초기 설정

   - **.env 파일 설정**  
     OpenWeather 및 OCR 서버 설정 정보를 `.env` 파일에 저장합니다.
     ```plaintext
     CURRENTAPI=[openweathermap API URL]
     KEY=[openweathermap API key]
     LANG="kr"
     OCRAPI=[OCR 서버 URL]
     GPTKEY=[gpt API key]
     ```

   - **채팅창 위치 설정**  
     `cmd/serve.go`의 `Start` 함수에서 다음과 같이 **messageBox**와 **speechBox**를 수정하여 스캔할 채팅 위치를 설정합니다.<br>
     **messageBox**는 채팅방 리스트화면입니다. 아래와 같이 빨간색알림이 뜨는 곳을 지정해야합니다.
     <img width="457" alt="image" src="https://github.com/user-attachments/assets/70a5cafd-62c1-42eb-9860-9ab23ab92e0d">
     
      **speechBox**는 채팅방 화면입니다. 아래와 같이 새로운 채팅이 뜨는 곳을 지정해야합니다.
     <img width="565" alt="image" src="https://github.com/user-attachments/assets/2186626d-4716-45f2-a86e-47e8bce32fb0">

     ```go
     func Start() {
         messageBox := &internal.DIR{X1: 490, Y1: 900, X2: 570, Y2: 955}
         speechBox := &internal.SpeechBox{X: 0, Y: 760, Width: 560, Height: 72}
     }
     ```

### 2. OpenWeather API 키 발급

   - [OpenWeather의 Current Weather API](https://openweathermap.org/current)에 가입하여 무료 API 키를 발급받습니다.

### 3. 지역 매칭 테이블 설정

   - 외국 API라 정확한 한국어 지역명을 위해 **매칭 테이블**을 설정합니다.
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
     ```
### 4. GPT API 키 발급
   - [OpenAI의 GPT API](https://platform.openai.com/api-keys)에 가입하여 API 키 발급합니다.
   - 필요한 경우 API사용을 위한 결제필요합니다.

### 5. OCR 서버 설정

   OCR 서버를 빌드하고 실행합니다.

   ```bash
   cd OCR
   docker build -t ocr .
   docker run -p 3000:5000 ocr
   ```

### 6. Bot 실행

   ```bash
   cd robot
   go run main.go serve
   ```

---

## 🛠️ 에러 해결

- **RobotGo 오류 발생 시**  
  [RobotGo](https://github.com/go-vgo/robotgo) 관련 오류가 발생할 경우, Xcode 설치가 필요할 수 있습니다.

---

## 🎥 실행 영상

- [영상] https://drive.google.com/file/d/16AKGU-Ld742PhNFq9rED9RfH4H4WXZv1/view?usp=sharing

---

## 📝 버전 정보

- **v1**: 지역 현재날씨 정보, gpt를 이용한 자동응답
