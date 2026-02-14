sequenceDiagram
    participant NextGenApp
    participant AuthBackend
    participant PWAWebView
    participant PWABackend

    NextGenApp->>Auth Backend: POST /api/pwa/auth/token<br/>{ msisdn, language, deviceId }
    Auth Backend->>NextGenApp: { accessToken JWT, sessionId UUID, expiresIn }
    NextGenApp->>PWAWebView: Load PWA with accessToken<br/>stored securely, not in URL if possible
    PWAWebView->>PWABackend: GET /api/session<br/>Authorization: Bearer <accessToken>
    PWABackend->>PWAWebView: Validate JWT<br/>Create session context<br/>Return session data
    PWAWebView->>NextGenApp: postMessage via JS Bridge<br/>type: "SESSION_INITIALIZED"<br/>sessionId: "<uuid>"
    Note over NextGenApp,PWABackend: Normal authenticated API calls continue<br/>Authorization: Bearer <accessToken>