// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovnnb

const HAChassisGroupTable = "HA_Chassis_Group"

// HAChassisGroup defines an object in HA_Chassis_Group table
type HAChassisGroup struct {
	UUID        string            `ovsdb:"_uuid"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	HaChassis   []string          `ovsdb:"ha_chassis"`
	Name        string            `ovsdb:"name"`
}
