# 🚀 Localhost Socket Relay Server

로컬 환경에서 동작하는 **소켓 릴레이 서버(Localhost Socket Relay Server)** 입니다.  
이 서버는 클라이언트 간의 소켓 통신을 중계(relay)하는 역할을 합니다.

---

## 🧩 설치 방법

### 1️⃣ Go 설치

이 프로그램은 **Go (Golang)** 으로 개발되었으므로, 먼저 Go를 설치해야 합니다.

🔗 설치 안내: [https://go.dev/doc/install](https://go.dev/doc/install)

설치가 완료되면 아래 명령어로 정상 설치 여부를 확인하세요.

```bash
go version
```

출력 예시:
```
go version go1.23.3 windows/amd64
```

---

### 2️⃣ 소스코드 빌드

다운로드받은 소스코드 폴더에서 마우스 오른쪽 클릭 후  
**“터미널에서 열기(T)”** 를 선택해 해당 경로에서 콘솔을 엽니다.

이후 다음 명령어를 실행합니다 👇

```bash
go build -o relay.exe
```

빌드가 완료되면 폴더 내에 `relay.exe` 파일이 생성됩니다.

---

### 3️⃣ 프로그램 실행

생성된 **`relay.exe`** 파일을 더블 클릭하거나,  
콘솔(Command Prompt)에서 직접 실행할 수 있습니다.

```bash
relay.exe
```

실행 후 서버가 정상적으로 시작되면 콘솔에 로그 메시지가 표시됩니다.

---

## ⚙️ 참고

- Go 환경 설정 시 **GOPATH** 및 **PATH** 설정이 되어 있어야 합니다.  
- 기본적으로 로컬호스트(`127.0.0.1`)에서만 작동합니다.  
- 방화벽 또는 포트 충돌이 있는 경우 정상적으로 실행되지 않을 수 있습니다.

---

## 🧠 예시 동작 구조 (개념도)

```
Client A  <---->  Relay Server  <---->  Client B
```

Relay Server가 중계 역할을 하여  
두 클라이언트가 직접 연결되지 않아도 데이터 송수신이 가능합니다.

---

## 📄 라이선스

이 프로젝트는 MIT License를 따릅니다.
