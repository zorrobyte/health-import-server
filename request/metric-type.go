package request

const MetricTypeQty = "QTY"
const MetricTypeMinMaxAvg = "MINMAXAVG"
const MetricTypeSleep = "SLEEP"
const MetricTypeUnknown = "UNKNOWN"

var metricTypeLookup = map[string]string{
	"active_energy":                      MetricTypeQty,
	"apple_exercise_time":                MetricTypeQty,
	"apple_stand_hour":                   MetricTypeQty,
	"apple_stand_time":                   MetricTypeQty,
	"basal_body_temperature":             MetricTypeUnknown,
	"basal_energy_burned":                MetricTypeQty,
	"blood_alcohol_content":              MetricTypeUnknown,
	"blood_glucose":                      MetricTypeUnknown,
	"blood_oxygen_saturation":            MetricTypeQty,
	"blood_pressure":                     MetricTypeUnknown,
	"body_fat_percentage":                MetricTypeUnknown,
	"body_mass_index":                    MetricTypeUnknown,
	"body_temperature":                   MetricTypeUnknown,
	"calcium":                            MetricTypeUnknown,
	"carbohydrates":                      MetricTypeUnknown,
	"chloride":                           MetricTypeUnknown,
	"chromium":                           MetricTypeUnknown,
	"copper":                             MetricTypeUnknown,
	"cycling_distance":                   MetricTypeUnknown,
	"dietary_biotin":                     MetricTypeUnknown,
	"dietary_caffeine":                   MetricTypeUnknown,
	"dietary_cholesterol":                MetricTypeUnknown,
	"dietary_energy":                     MetricTypeUnknown,
	"dietary_sugar":                      MetricTypeUnknown,
	"dietary_water":                      MetricTypeUnknown,
	"distance_downhill_snow_sports":      MetricTypeUnknown,
	"environmental_audio_exposure":       MetricTypeQty,
	"fiber":                              MetricTypeUnknown,
	"flights_climbed":                    MetricTypeQty,
	"folate":                             MetricTypeUnknown,
	"forced_expiratory_volume_1":         MetricTypeUnknown,
	"forced_vital_capacity":              MetricTypeUnknown,
	"handwashing":                        MetricTypeQty,
	"headphone_audio_exposure":           MetricTypeQty,
	"heart_rate":                         MetricTypeMinMaxAvg,
	"heart_rate_variability":             MetricTypeQty,
	"height":                             MetricTypeQty,
	"high_heart_rate_notifications":      MetricTypeUnknown,
	"inhaler_usage":                      MetricTypeUnknown,
	"insulin_delivery":                   MetricTypeUnknown,
	"iodine":                             MetricTypeUnknown,
	"iron":                               MetricTypeUnknown,
	"irregular_heart_rate_notifications": MetricTypeUnknown,
	"lean_body_mass":                     MetricTypeUnknown,
	"low_heart_rate_notifications":       MetricTypeUnknown,
	"magnesium":                          MetricTypeUnknown,
	"manganese":                          MetricTypeUnknown,
	"mindful_minutes":                    MetricTypeQty,
	"molybdenum":                         MetricTypeUnknown,
	"monounsaturated_fat":                MetricTypeUnknown,
	"niacin":                             MetricTypeUnknown,
	"number_of_times_fallen":             MetricTypeUnknown,
	"pantothenic_acid":                   MetricTypeUnknown,
	"peripheral_perfusion_index":         MetricTypeUnknown,
	"polyunsaturated_fat":                MetricTypeUnknown,
	"potassium":                          MetricTypeUnknown,
	"protein":                            MetricTypeUnknown,
	"push_count":                         MetricTypeUnknown,
	"respiratory_rate":                   MetricTypeUnknown,
	"resting_heart_rate":                 MetricTypeQty,
	"riboflavin":                         MetricTypeUnknown,
	"saturated_fat":                      MetricTypeUnknown,
	"selenium":                           MetricTypeUnknown,
	"sexual_activity":                    MetricTypeUnknown,
	"six-minute_walking_test_distance":   MetricTypeQty,
	"sleep_analysis":                     MetricTypeSleep,
	"sodium":                             MetricTypeUnknown,
	"stair_speed:_down":                  MetricTypeQty,
	"stair_speed:_up":                    MetricTypeQty,
	"step_count":                         MetricTypeQty,
	"swimming_distance":                  MetricTypeUnknown,
	"swimming_stroke_count":              MetricTypeUnknown,
	"thiamin":                            MetricTypeUnknown,
	"toothbrushing":                      MetricTypeUnknown,
	"total_fat":                          MetricTypeUnknown,
	"vo2_max":                            MetricTypeQty,
	"vitamin_a":                          MetricTypeUnknown,
	"vitamin_b12":                        MetricTypeUnknown,
	"vitamin_b6":                         MetricTypeUnknown,
	"vitamin_c":                          MetricTypeUnknown,
	"vitamin_d":                          MetricTypeUnknown,
	"vitamin_e":                          MetricTypeUnknown,
	"vitamin_k":                          MetricTypeUnknown,
	"waist_circumference":                MetricTypeUnknown,
	"walking_running_distance":           MetricTypeQty,
	"walking_asymmetry_percentage":       MetricTypeQty,
	"walking_double_support_percentage":  MetricTypeQty,
	"walking_heart_rate_average":         MetricTypeQty,
	"walking_speed":                      MetricTypeQty,
	"walking_step_length":                MetricTypeQty,
	"weight_body_mass":                   MetricTypeQty,
	"wheelchair_distance":                MetricTypeUnknown,
	"zinc":                               MetricTypeUnknown,
}

func LookupMetricType(metricName string) string {
	metricType, ok := metricTypeLookup[metricName]
	if ok {
		return metricType
	}
	return MetricTypeUnknown
}
