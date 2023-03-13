package model
import (
    "testing"
)

func TestNewOriginCorrectConnection(t *testing.T) {
    domain, _ := NewDomain(1, "test_domain")
    connInfo := map[string]string{
        "key": "string",
    }
    origin, _ := NewOrigin(2, "test_origin", connInfo)

    connVals := ConnectionValues{
        "key": "test_key",
    }

    _, err := NewOriginInstance(0, "test_instance",
                                origin, domain, connVals)
    if err != nil {
        t.Fatalf("Test failure: NewOriginInstance returned error %v", err)
    }
}

func TestNewOriginInCorrectConnection(t *testing.T) {
    domain, _ := NewDomain(1, "test_domain")
    connInfo := map[string]string{
        "key": "string",
    }
    origin, _ := NewOrigin(2, "test_origin", connInfo)

    connVals := ConnectionValues{
        "wrongkey": "test_key",
    }

    _, err := NewOriginInstance(0, "test_instance",
                                origin, domain, connVals)
    if err == nil {
        t.Fatalf("Test failure: NewOriginInstance did not "+
                 "detect incorrect connection info format")
    }
}
