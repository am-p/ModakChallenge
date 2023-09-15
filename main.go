package main

import (
    "fmt"
    "time"
)

type NotificationService struct {
    updateEmailCh   chan string
    newsEmailCh     chan string
    marketingEmailCh chan string
}

func NewNotificationService() *NotificationService {
    return &NotificationService{
        updateEmailCh:   make(chan string, 2),
        newsEmailCh:     make(chan string, 1),
        marketingEmailCh: make(chan string, 1),
    }
}

func (ns *NotificationService) SendUpdateEmail(recipient string) error {
    select {
    case ns.updateEmailCh <- recipient:
        return nil
    default:
        return fmt.Errorf("Rate limit exceeded for update email")
    }
}

func (ns *NotificationService) SendNewsEmail(recipient string) error {
    select {
    case ns.newsEmailCh <- recipient:
        return nil
    default:
        return fmt.Errorf("Rate limit exceeded for news email")
    }
}

func (ns *NotificationService) SendMarketingEmail(recipient string) error {
    select {
    case ns.marketingEmailCh <- recipient:
        return nil
    default:
        return fmt.Errorf("Rate limit exceeded for marketing email")
    }
}

func main() {
    ns := NewNotificationService()
    recipient := "ariel@modak.com"
    for i := 0; i < 10; i++ {
        err := ns.SendUpdateEmail(recipient)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Update email sent successfully")
        }

        err = ns.SendNewsEmail(recipient)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("News email sent successfully")
        }

        err = ns.SendMarketingEmail(recipient)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println("Marketing email sent successfully")
        }

        time.Sleep(30 * time.Second)
    }
}
