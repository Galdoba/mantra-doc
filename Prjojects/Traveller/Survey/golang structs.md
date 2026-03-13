---
updated_at: 2026-03-13T14:22:37.973+10:00
---
```go
package iissforms

// StarSystem представляет полные данные о звёздной системе.
// Может быть сконвертирован в любую из форм IISS с помощью соответствующих методов.
type StarSystem struct {
	// Общая информация (присутствует во многих формах)
	SectorGrid      string  `json:"sector_grid,omitempty"`
	InitialSurvey   string  `json:"initial_survey,omitempty"`
	LastUpdated     string  `json:"last_updated,omitempty"`
	IISSDesignation string  `json:"iiss_designation,omitempty"`
	SystemAge       float64 `json:"system_age_gyr,omitempty"`
	TravelZone      string  `json:"travel_zone,omitempty"`

	// Звёзды
	Stars []Star `json:"stars,omitempty"`

	// Планеты (включая газовые гиганты, планетезимали и т.д.)
	Planets []Planet `json:"planets,omitempty"`

	// Пояса астероидов
	Belts []Belt `json:"belts,omitempty"`

	// Дополнительная информация о системе (например, из Class 0 Sector Survey)
	StellarObjectsSummary *StellarObjectsSummary `json:"stellar_objects_summary,omitempty"`

	// Комментарии (общие)
	Comments string `json:"comments,omitempty"`
}

// Star представляет звезду в системе.
type Star struct {
	Designation  string  `json:"designation,omitempty"`   // например, "Aa"
	Component    string  `json:"component,omitempty"`     // "Aa", "Ab"...
	Class        string  `json:"class,omitempty"`         // "G7 V"
	Mass         float64 `json:"mass,omitempty"`          // в солнечных массах
	Temperature  int     `json:"temperature,omitempty"`   // в K
	Diameter     float64 `json:"diameter,omitempty"`      // в солнечных диаметрах
	Luminosity   float64 `json:"luminosity,omitempty"`    // в солнечных светимостях
	OrbitNum     float64 `json:"orbit_number,omitempty"`
	AU           float64 `json:"au,omitempty"`
	Eccentricity float64 `json:"eccentricity,omitempty"`
	Period       string  `json:"period,omitempty"` // в годах или днях
	HZCO         float64 `json:"hzco,omitempty"`   // зона обитаемости
	MAO          string  `json:"mao,omitempty"`    // из Class II/III
	Notes        string  `json:"notes,omitempty"`
}

// Planet представляет планету или крупное тело.
type Planet struct {
	Name        string `json:"name,omitempty"`
	Designation string `json:"designation,omitempty"` // например, "Ab"
	Primary     string `json:"primary,omitempty"`     // звезда, вокруг которой вращается
	UWP         string `json:"uwp,omitempty"`         // Universal World Profile
	SAHUWP      string `json:"sah_uwp,omitempty"`     // другой формат UWP (из форм)
	OrbitNum    int    `json:"orbit_number,omitempty"`
	AU          float64 `json:"au,omitempty"`
	Eccentricity float64 `json:"eccentricity,omitempty"`
	Period      string  `json:"period,omitempty"`
	Type        string  `json:"type,omitempty"`       // "Gas Giant", "Terrestrial", etc.
	Diameter    float64 `json:"diameter,omitempty"`   // в км
	Density     float64 `json:"density,omitempty"`    // в г/см³
	Mass        float64 `json:"mass,omitempty"`       // в массах Земли или кг?

	// Детальные характеристики (из Form0407KIVPartP)
	Size         *PlanetSize         `json:"size,omitempty"`
	Atmosphere   *PlanetAtmosphere   `json:"atmosphere,omitempty"`
	Hydrographics *PlanetHydrographics `json:"hydrographics,omitempty"`
	Rotation     *PlanetRotation     `json:"rotation,omitempty"`
	Temperature  *PlanetTemperature  `json:"temperature,omitempty"`
	Life         *PlanetLife         `json:"life,omitempty"`
	Resources    *ResourcesInfo      `json:"resources,omitempty"`
	Habitability *HabitabilityInfo   `json:"habitability,omitempty"`

	// Социально-экономические характеристики (из Form0407FIVPartC)
	Population  *PopulationInfo `json:"population,omitempty"`
	Government  *GovernmentInfo `json:"government,omitempty"`
	LawLevel    *LawLevelInfo   `json:"law_level,omitempty"`
	Technology  *TechnologyInfo `json:"technology,omitempty"`
	Culture     *CultureInfo    `json:"culture,omitempty"`
	Economics   *EconomicsInfo  `json:"economics,omitempty"`
	Starport    *StarportInfo   `json:"starport,omitempty"`
	Military    *MilitaryInfo   `json:"military,omitempty"`

	// Спутники
	Satellites []Satellite `json:"satellites,omitempty"`

	Notes string `json:"notes,omitempty"`
}

// Belt представляет пояс астероидов.
type Belt struct {
	Designation    string          `json:"designation,omitempty"`
	Primary        string          `json:"primary,omitempty"` // звезда
	OrbitNum       int             `json:"orbit_number,omitempty"`
	AU             float64         `json:"au,omitempty"`
	Eccentricity   float64         `json:"eccentricity,omitempty"`
	Period         string          `json:"period,omitempty"`
	Composition    *BeltComposition `json:"composition,omitempty"`
	MajorBodiesSize1 []string       `json:"size_1_bodies,omitempty"`
	MajorBodiesSizeS []string       `json:"size_s_bodies,omitempty"`
	Notes          string           `json:"notes,omitempty"`
}

// BeltComposition описывает состав пояса.
type BeltComposition struct {
	MTypePercent string `json:"m_type_percent,omitempty"`
	STypePercent string `json:"s_type_percent,omitempty"`
	CTypePercent string `json:"c_type_percent,omitempty"`
	OtherPercent string `json:"other_percent,omitempty"`
	Bulk         string `json:"bulk,omitempty"`
}

// Satellite представляет спутник планеты.
type Satellite struct {
	Designation  string  `json:"designation,omitempty"`
	SAHUWP       string  `json:"sah_uwp,omitempty"`
	OrbitPD      string  `json:"orbit_pd,omitempty"` // в диаметрах планеты
	OrbitKm      float64 `json:"orbit_km,omitempty"`
	Eccentricity float64 `json:"eccentricity,omitempty"`
	Diameter     float64 `json:"diameter,omitempty"` // в км
	Density      float64 `json:"density,omitempty"`
	Mass         float64 `json:"mass,omitempty"`
	PeriodHours  float64 `json:"period_hours,omitempty"`
	SizeDeg      float64 `json:"size_deg,omitempty"` // видимый размер в градусах
	Notes        string  `json:"notes,omitempty"`
}

// --- Детальные структуры для планет (аналоги Form0407KIVPartP) ---

type PlanetSize struct {
	Diameter       float64 `json:"diameter,omitempty"`
	Composition    string  `json:"composition,omitempty"`
	Density        float64 `json:"density,omitempty"`
	Gravity        float64 `json:"gravity,omitempty"`
	Mass           float64 `json:"mass,omitempty"`
	EscapeVelocity string  `json:"escape_velocity_kps,omitempty"`
	Notes          string  `json:"notes,omitempty"`
}

type PlanetAtmosphere struct {
	Pressure    float64 `json:"pressure_bar,omitempty"`
	Composition string  `json:"composition,omitempty"`
	Oxygen      float64 `json:"oxygen_bar,omitempty"`
	Taints      string  `json:"taints,omitempty"`
	ScaleHeight string  `json:"scale_height,omitempty"`
	Notes       string  `json:"notes,omitempty"`
}

type PlanetHydrographics struct {
	Coverage     string `json:"coverage_percent,omitempty"`
	Composition  string `json:"composition,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	MajorBodies  string `json:"major_bodies,omitempty"`
	MinorBodies  string `json:"minor_bodies,omitempty"`
	Other        string `json:"other,omitempty"`
	Notes        string `json:"notes,omitempty"`
}

type PlanetRotation struct {
	Sidereal      string  `json:"sidereal,omitempty"`
	Solar         string  `json:"solar,omitempty"`
	SolarDaysYear float64 `json:"solar_days_per_year,omitempty"`
	AxialTilt     string  `json:"axial_tilt,omitempty"`
	TidalLock     bool    `json:"tidal_lock,omitempty"`
	Tides         string  `json:"tides,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}

type PlanetTemperature struct {
	High            string  `json:"high,omitempty"`
	Mean            string  `json:"mean,omitempty"`
	Low             string  `json:"low,omitempty"`
	Luminosity      float64 `json:"luminosity,omitempty"`
	Albedo          float64 `json:"albedo,omitempty"`
	Greenhouse      float64 `json:"greenhouse,omitempty"`
	Notes           string  `json:"notes,omitempty"`
	SeismicStress   string  `json:"seismic_stress,omitempty"`
	ResidualStress  string  `json:"residual_stress,omitempty"`
	TidalStress     string  `json:"tidal_stress,omitempty"`
	TidalHeating    string  `json:"tidal_heating,omitempty"`
	TectonicPlates  string  `json:"major_tectonic_plates,omitempty"`
}

type PlanetLife struct {
	Biomass        string `json:"biomass,omitempty"`
	Biocomplexity  string `json:"biocomplexity,omitempty"`
	Sophonts       string `json:"sophonts,omitempty"`
	Biodiversity   string `json:"biodiversity,omitempty"`
	Compatibility  string `json:"compatibility,omitempty"`
	Notes          string `json:"notes,omitempty"`
}

type ResourcesInfo struct {
	Rating string `json:"rating,omitempty"`
	Notes  string `json:"notes,omitempty"`
}

type HabitabilityInfo struct {
	Rating string `json:"rating,omitempty"`
	Notes  string `json:"notes,omitempty"`
}

// --- Социально-экономические структуры (аналоги Form0407FIVPartC) ---

type PopulationInfo struct {
	Total          string   `json:"total,omitempty"`
	Demographics   string   `json:"demographics,omitempty"`
	PCR            string   `json:"pcr,omitempty"`
	Urbanisation   string   `json:"urbanisation_percent,omitempty"`
	MajorCities    int      `json:"major_cities,omitempty"`
	CapitalPort    string   `json:"capital_port,omitempty"`
	MajorCitiesList []string `json:"major_cities_list,omitempty"`
	Notes          string   `json:"notes,omitempty"`
}

type GovernmentInfo struct {
	Type           string        `json:"type,omitempty"`
	Centralization string        `json:"centralization,omitempty"`
	Authority      string        `json:"authority,omitempty"`
	Profile        string        `json:"profile,omitempty"`
	Notes          string        `json:"notes,omitempty"`
	Factions       []FactionInfo `json:"factions,omitempty"`
}

type FactionInfo struct {
	Name          string `json:"name,omitempty"`
	Profile       string `json:"profile,omitempty"`
	Description   string `json:"description,omitempty"`
	Relationships string `json:"relationships,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

type LawLevelInfo struct {
	Primary               string        `json:"primary,omitempty"`
	Secondary             string        `json:"secondary,omitempty"`
	Uniformity            string        `json:"uniformity,omitempty"`
	PresumptionOfInnocence bool          `json:"presumption_of_innocence,omitempty"`
	DeathPenalty          bool          `json:"death_penalty,omitempty"`
	Categories            *LawCategories `json:"categories,omitempty"`
	Notes                 string        `json:"notes,omitempty"`
}

type LawCategories struct {
	Overall       string `json:"overall,omitempty"`
	Weapons       string `json:"weapons,omitempty"`
	Economics     string `json:"economics,omitempty"`
	Criminal      string `json:"criminal,omitempty"`
	Private       string `json:"private,omitempty"`
	PersonalRights string `json:"personal_rights,omitempty"`
}

type TechnologyInfo struct {
	CommonHigh      string `json:"common_high,omitempty"`
	CommonLow       string `json:"common_low,omitempty"`
	Energy          string `json:"energy,omitempty"`
	Electronics     string `json:"electronics,omitempty"`
	Manufacturing   string `json:"manufacturing,omitempty"`
	Medical         string `json:"medical,omitempty"`
	Land            string `json:"land,omitempty"`
	Water           string `json:"water,omitempty"`
	Air             string `json:"air,omitempty"`
	Space           string `json:"space,omitempty"`
	PersonalMilitary string `json:"personal_military,omitempty"`
	HeavyMilitary   string `json:"heavy_military,omitempty"`
	Novelty         string `json:"novelty,omitempty"`
	Environmental   string `json:"environmental,omitempty"`
	Notes           string `json:"notes,omitempty"`
}

type CultureInfo struct {
	Diversity      string `json:"diversity,omitempty"`
	Cohesion       string `json:"cohesion,omitempty"`
	Xenophilia     string `json:"xenophilia,omitempty"`
	Progressiveness string `json:"progressiveness,omitempty"`
	Uniqueness     string `json:"uniqueness,omitempty"`
	Expansionism   string `json:"expansionism,omitempty"`
	Symbology      string `json:"symbology,omitempty"`
	Militancy      string `json:"militancy,omitempty"`
	Notes          string `json:"notes,omitempty"`
}

type EconomicsInfo struct {
	TradeCodes      string  `json:"trade_codes,omitempty"`
	Importance      string  `json:"importance,omitempty"`
	Resources       string  `json:"resources,omitempty"`
	Labour          string  `json:"labour,omitempty"`
	Infrastructure  string  `json:"infrastructure,omitempty"`
	Efficiency      string  `json:"efficiency,omitempty"`
	GWPPerCapita    float64 `json:"gwp_per_capita,omitempty"`
	WTN             float64 `json:"wtn,omitempty"`
	InequalityRating string  `json:"inequality_rating,omitempty"`
	DevelopmentScore string  `json:"development_score,omitempty"`
	GWP             float64 `json:"gwp_mcr,omitempty"`
	Tariffs         string  `json:"tariffs,omitempty"`
	Notes           string  `json:"notes,omitempty"`
}

type StarportInfo struct {
	Class               string `json:"class,omitempty"`
	Highport            string `json:"highport,omitempty"`
	ExpectedWeeklyTraffic string `json:"expected_weekly_traffic,omitempty"`
	BerthingFees        string `json:"berthing_fees,omitempty"`
	Docking             string `json:"docking,omitempty"`
	Shipyard            string `json:"shipyard,omitempty"`
	AnnualOutput        string `json:"annual_output,omitempty"`
	NavyBase            string `json:"navy_base,omitempty"`
	ScoutBase           string `json:"scout_base,omitempty"`
	MilitaryBase        string `json:"military_base,omitempty"`
	Other               string `json:"other,omitempty"`
	Notes               string `json:"notes,omitempty"`
}

type MilitaryInfo struct {
	EffectiveBudgetPercent string `json:"effective_budget_percent,omitempty"`
	Structure              string `json:"structure,omitempty"`
	Enforcement            string `json:"enforcement,omitempty"`
	Militia                string `json:"militia,omitempty"`
	Army                   string `json:"army,omitempty"`
	WetNavy                string `json:"wet_navy,omitempty"`
	AirForce               string `json:"air_force,omitempty"`
	SystemDefence          string `json:"system_defence,omitempty"`
	Navy                   string `json:"navy,omitempty"`
	Marines                string `json:"marines,omitempty"`
	Notes                  string `json:"notes,omitempty"`
}

// StellarObjectsSummary соответствует полю "Stellar objects" из Form0398D0.
type StellarObjectsSummary struct {
	Location string `json:"location,omitempty"`
	Primary  string `json:"primary,omitempty"`
	PrimaryP string `json:"primary_plus,omitempty"`
	Close    string `json:"close,omitempty"`
	CloseP   string `json:"close_plus,omitempty"`
	Near     string `json:"near,omitempty"`
	NearP    string `json:"near_plus,omitempty"`
	Far      string `json:"far,omitempty"`
	FarP     string `json:"far_plus,omitempty"`
	GG       int    `json:"gg,omitempty"`
	Notes    string `json:"notes,omitempty"`
}
```
