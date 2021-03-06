package awx

import (
	"testing"
	"time"
)

func TestListGroups(t *testing.T) {
	var (
		expectListGroupsResponse = []*Group{
			{
				ID:   21,
				Type: 0,
				URL:  "/api/v2/groups/21/",
				Related: &Related{
					CreatedBy:         "/api/v2/users/11/",
					JobHostSummaries:  "/api/v2/groups/21/job_host_summaries/",
					VariableData:      "/api/v2/groups/21/variable_data/",
					JobEvents:         "/api/v2/groups/21/job_events/",
					PotentialChildren: "/api/v2/groups/21/potential_children/",
					AdHocCommands:     "/api/v2/groups/21/ad_hoc_commands/",
					AllHosts:          "/api/v2/groups/21/all_hosts/",
					ActivityStream:    "/api/v2/groups/21/activity_stream/",
					Hosts:             "/api/v2/groups/21/hosts/",
					Children:          "/api/v2/groups/21/children/",
					InventorySources:  "/api/v2/groups/21/inventory_sources/",
					Inventory:         "/api/v2/inventories/9/",
				},
				SummaryFields: &Summary{
					Inventory: &Inventory{
						ID:                           9,
						Name:                         "test",
						Description:                  "",
						HasActiveFailures:            false,
						TotalHosts:                   8,
						HostsWithActiveFailures:      0,
						TotalGroups:                  4,
						GroupsWithActiveFailures:     0,
						HasInventorySources:          false,
						TotalInventorySources:        0,
						InventorySourcesWithFailures: 0,
						OrganizationID:               16,
						Kind:                         "",
					},
					CreatedBy: &ByUserSummary{
						ID:        11,
						Username:  "demouser",
						FirstName: "Demo",
						LastName:  "User",
					},
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Copy:   true,
						Delete: true,
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-07-17T13:27:46.686176Z")
					return t
				}(),
				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2018-07-17T13:28:07.127040Z")
					return t
				}(),
				Name:                     "Demo Group",
				Description:              "",
				Inventory:                9,
				Variables:                "",
				HasActiveFailures:        false,
				TotalHosts:               3,
				HostsWithActiveFailures:  0,
				TotalGroups:              0,
				GroupsWithActiveFailures: 0,
				HasInventorySources:      false,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.GroupService.ListGroups(map[string]string{
		"name": "Demo Group",
	})

	if err != nil {
		t.Fatalf("ListGroups err: %s", err)
	} else {
		checkAPICallResult(t, expectListGroupsResponse, result)
		t.Log("ListGroups passed!")
	}
}

func TestCreateGroup(t *testing.T) {
	var (
		expectCreateGroupResponse = &Group{
			ID:   21,
			Type: 0,
			URL:  "/api/v2/groups/21/",
			Related: &Related{
				CreatedBy:         "/api/v2/users/11/",
				JobHostSummaries:  "/api/v2/groups/21/job_host_summaries/",
				VariableData:      "/api/v2/groups/21/variable_data/",
				JobEvents:         "/api/v2/groups/21/job_events/",
				PotentialChildren: "/api/v2/groups/21/potential_children/",
				AdHocCommands:     "/api/v2/groups/21/ad_hoc_commands/",
				AllHosts:          "/api/v2/groups/21/all_hosts/",
				ActivityStream:    "/api/v2/groups/21/activity_stream/",
				Hosts:             "/api/v2/groups/21/hosts/",
				Children:          "/api/v2/groups/21/children/",
				InventorySources:  "/api/v2/groups/21/inventory_sources/",
				Inventory:         "/api/v2/inventories/9/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           9,
					Name:                         "test",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   8,
					HostsWithActiveFailures:      0,
					TotalGroups:                  4,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               16,
					Kind:                         "",
				},
				CreatedBy: &ByUserSummary{
					ID:        11,
					Username:  "demouser",
					FirstName: "Demo",
					LastName:  "User",
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Delete: true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:27:46.686176Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:28:07.127040Z")
				return t
			}(),
			Name:                     "Demo Group",
			Description:              "",
			Inventory:                9,
			Variables:                "",
			HasActiveFailures:        false,
			TotalHosts:               3,
			HostsWithActiveFailures:  0,
			TotalGroups:              0,
			GroupsWithActiveFailures: 0,
			HasInventorySources:      false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.CreateGroup(map[string]interface{}{
		"name":        "Test Group",
		"description": "for testing CreateGroup api",
		"inventory":   1,
		"variables":   "",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("CreateGroup err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateGroupResponse, result)
		t.Log("CreateGroup passed!")
	}
}

func TestUpdateGroup(t *testing.T) {
	var (
		expectUpdateGroupResponse = &Group{
			ID:   21,
			Type: 0,
			URL:  "/api/v2/groups/21/",
			Related: &Related{
				CreatedBy:         "/api/v2/users/11/",
				JobHostSummaries:  "/api/v2/groups/21/job_host_summaries/",
				VariableData:      "/api/v2/groups/21/variable_data/",
				JobEvents:         "/api/v2/groups/21/job_events/",
				PotentialChildren: "/api/v2/groups/21/potential_children/",
				AdHocCommands:     "/api/v2/groups/21/ad_hoc_commands/",
				AllHosts:          "/api/v2/groups/21/all_hosts/",
				ActivityStream:    "/api/v2/groups/21/activity_stream/",
				Hosts:             "/api/v2/groups/21/hosts/",
				Children:          "/api/v2/groups/21/children/",
				InventorySources:  "/api/v2/groups/21/inventory_sources/",
				Inventory:         "/api/v2/inventories/9/",
			},
			SummaryFields: &Summary{
				Inventory: &Inventory{
					ID:                           9,
					Name:                         "test",
					Description:                  "",
					HasActiveFailures:            false,
					TotalHosts:                   8,
					HostsWithActiveFailures:      0,
					TotalGroups:                  4,
					GroupsWithActiveFailures:     0,
					HasInventorySources:          false,
					TotalInventorySources:        0,
					InventorySourcesWithFailures: 0,
					OrganizationID:               16,
					Kind:                         "",
				},
				CreatedBy: &ByUserSummary{
					ID:        11,
					Username:  "demouser",
					FirstName: "Demo",
					LastName:  "User",
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Copy:   true,
					Delete: true,
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:27:46.686176Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2018-07-17T13:28:07.127040Z")
				return t
			}(),
			Name:                     "Demo Group",
			Description:              "Add description",
			Inventory:                9,
			Variables:                "",
			HasActiveFailures:        false,
			TotalHosts:               3,
			HostsWithActiveFailures:  0,
			TotalGroups:              0,
			GroupsWithActiveFailures: 0,
			HasInventorySources:      false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.UpdateGroup(21, map[string]interface{}{
		"description": "Add description",
	}, map[string]string{})

	if err != nil {
		t.Fatalf("UpdateGroup err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateGroupResponse, result)
		t.Log("UpdateGroup passed!")
	}
}
func TestDeleteGroup(t *testing.T) {
	var (
		expectDeleteGroupResponse = &Group{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.GroupService.DeleteGroup(21)

	if err != nil {
		t.Fatalf("DeleteGroup err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteGroupResponse, result)
		t.Log("DeleteGroup passed!")
	}
}
