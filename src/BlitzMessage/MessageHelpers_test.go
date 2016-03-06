//  HappinessServer  -  Track the Happiness user data.
//
//  E.B.Smith  -  November, 2014


package happiness

import (
    "testing"
    "database/sql"
    "github.com/golang/protobuf/proto"
)


func StringPtr(s string) *string {
    return &s
}
func Float64Ptr(f float64) *float64 {
    return &f
}


func Test_NullStringFromScoreComponents(t *testing.T) {

    components := [2]*ScoreComponent {
        &ScoreComponent{ Label: StringPtr("team"),  Score: Float64Ptr(1.0) },
        &ScoreComponent{ Label: StringPtr("trust"), Score: Float64Ptr(2.0) },
    }

    ns := NullStringFromScoreComponents(components[:])

    if ! ns.Valid {
        t.Errorf("Got not valid.")
        return
    }

    s := " { \"(team, 1.000000)\", \"(trust, 2.000000)\" } "
    if ns.String != s {
        t.Errorf("Got '%s' but wanted '%s'.", ns.String, s)
        return
    }
}


func Test_ScoreComponentsFromNullString(t *testing.T) {
    //   {"(team,0.4)","(trust,0.26)"}

    s := sql.NullString {
        Valid:  true,
        String: "{\"(team,0.4)\",\"(trust,0.26)\"}",
    }

    components := ScoreComponentsFromNullString(s)
    if len(components) != 2 {
        t.Errorf("Failed.  Got: %+v", components)
        return
    }
    if *components[0].Label != "team" || *components[0].Score != 0.4 {
        t.Errorf("Failed.  Got: %+v", components)
        return
    }
    if *components[1].Label != "trust" || *components[1].Score != 0.26 {
        t.Errorf("Failed.  Got: %+v", components)
        return
    }
}


func Test_ClientRequest(t *testing.T) {

    userIDs := []string { "9876" }
    profileQuery := ProfileQuery { UserIDs: userIDs }
    request := ClientRequest { }
    request.SessionToken = StringPtr("1234")
    request.ClientRequestMessage = &ClientRequest_ProfileQuery { ProfileQuery: &profileQuery }

    t.Logf("    Request: %+v.", request)
    t.Logf("    Message: %+v.", request.GetProfileQuery().UserIDs)
    data, error := proto.Marshal(&request)
    if error != nil {
        t.Errorf("Marshal failed: %v.", error)
        return
    }

    newRequest := ClientRequest {}
    error = proto.Unmarshal(data, &newRequest)
    if error != nil {
        t.Errorf("Unmarshal failed: %v.", error)
        return
    }
    t.Logf("New request: %+v.", newRequest)

    if *newRequest.SessionToken != "1234" {
        t.Errorf("Session failed.")
        return
    }

    query := newRequest.GetProfileQuery()
    if query == nil || len(query.UserIDs) != 1 || query.UserIDs[0] != "9876" {
        t.Errorf("Wrong profile query: %+v.", newRequest)
    }

    // switch clientMessage := newRequest.(type) {
    // case ClientRequest_ProfileQuery:
    //     if clientMessage.UserIDs[0] != "9876" {
    //         t.Errorf("Wrong profile query.")
    //     }
    // default:
    //     t.Errorf("Wrong profile query.")
    // }
}

