package email

import (
    "fmt"
    "net/smtp"
    "os"
)

type EmailService struct {
    host     string
    port     string
    username string
    password string
    from     string
}

type EmailData struct {
    To      string
    Subject string
    Body    string
    IsHTML  bool
}

func NewEmailService() *EmailService {
    return &EmailService{
        host:     getEnv("SMTP_HOST", "smtp.gmail.com"),
        port:     getEnv("SMTP_PORT", "587"),
        username: getEnv("SMTP_USER", ""),
        password: getEnv("SMTP_PASSWORD", ""),
        from:     getEnv("SMTP_FROM", "noreply@zafood.com"),
    }
}

func (e *EmailService) SendEmail(data EmailData) error {
    if e.username == "" || e.password == "" {
        // Skip if no credentials (demo mode)
        fmt.Println("Demo mode: Email would be sent to", data.To)
        fmt.Println("Subject:", data.Subject)
        fmt.Println("Body:", data.Body)
        return nil
    }

    auth := smtp.PlainAuth("", e.username, e.password, e.host)
    
    to := []string{data.To}
    
    msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", 
        data.To, data.Subject, data.Body))
    
    addr := fmt.Sprintf("%s:%s", e.host, e.port)
    return smtp.SendMail(addr, auth, e.from, to, msg)
}

func (e *EmailService) SendOrderConfirmation(to string, orderID string, total float64) error {
    subject := fmt.Sprintf("Order Confirmation #%s", orderID)
    body := fmt.Sprintf(`
Hello!

Your order #%s has been confirmed!
Total amount: %.2f ₸
Estimated delivery: 30-45 minutes

Thank you for choosing ZAfood!

Best regards,
ZAfood Team
`, orderID, total)
    
    return e.SendEmail(EmailData{
        To:      to,
        Subject: subject,
        Body:    body,
        IsHTML:  false,
    })
}

func (e *EmailService) SendOrderStatusUpdate(to string, orderID string, status string) error {
    subject := fmt.Sprintf("Order #%s Status Update", orderID)
    body := fmt.Sprintf(`
Your order #%s status has been updated to: %s

Track your order in real-time on our website.

Best regards,
ZAfood Team
`, orderID, status)
    
    return e.SendEmail(EmailData{
        To:      to,
        Subject: subject,
        Body:    body,
        IsHTML:  false,
    })
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
