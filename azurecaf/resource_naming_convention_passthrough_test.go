package azurecaf

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccCafNamingConventionPassthrough(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePassthroughConfig,
				Check: resource.ComposeTestCheckFunc(

					testAccCafNamingValidation(
						"azurecaf_naming_convention.logs_inv",
						"logsinvalid",
						11,
						"log"),
					regexMatch("azurecaf_naming_convention.logs_inv", regexp.MustCompile(Resources["la"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_agw",
						"TEST-DEV-AGW-RG",
						15,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_agw", regexp.MustCompile(Resources["agw"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_apim",
						"TESTDEVAPIMRG",
						13,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_apim", regexp.MustCompile(Resources["apim"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_app",
						"TEST-DEV-APP-RG",
						15,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_app", regexp.MustCompile(Resources["app"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_appi",
						"TEST-DEV-APPI-RG",
						16,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_appi", regexp.MustCompile(Resources["appi"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_ase",
						"TEST-DEV-ASE-RG",
						15,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_ase", regexp.MustCompile(Resources["ase"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_plan",
						"TEST-DEV-PLAN-RG",
						16,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_plan", regexp.MustCompile(Resources["plan"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_sql",
						"test-dev-sql-rg",
						15,
						"test"),
					regexMatch("azurecaf_naming_convention.passthrough_sql", regexp.MustCompile(Resources["sql"].ValidationRegExp), 1),
					testAccCafNamingValidation(
						"azurecaf_naming_convention.passthrough_sqldb",
						"TEST-DEV-SQLDB-RG",
						17,
						"TEST"),
					regexMatch("azurecaf_naming_convention.passthrough_sqldb", regexp.MustCompile(Resources["sqldb"].ValidationRegExp), 1),
				),
			},
		},
	})
}

const testAccResourcePassthroughConfig = `
provider "azurecaf" {

}

#Storage account test
resource "azurecaf_naming_convention" "logs_inv" {
    convention      = "passthrough"
    name            = "logs_invalid"
    resource_type   = "la"
}

# Application Gateway
resource "azurecaf_naming_convention" "passthrough_agw" {
    convention      = "passthrough"
    name            = "TEST-DEV-AGW-RG"
    resource_type   = "azurerm_application_gateway"
}

# API Management
resource "azurecaf_naming_convention" "passthrough_apim" {
    convention      = "passthrough"
    name            = "TEST-DEV-APIM-RG"
    resource_type   = "azurerm_api_management"
}

# App Service
resource "azurecaf_naming_convention" "passthrough_app" {
    convention      = "passthrough"
    name            = "TEST-DEV-APP-RG"
    resource_type   = "azurerm_app_service"
}

# Application Insights
resource "azurecaf_naming_convention" "passthrough_appi" {
    convention      = "passthrough"
    name            = "TEST-DEV-APPI-RG"
    resource_type   = "azurerm_application_insights"
}

# App Service Environment
resource "azurecaf_naming_convention" "passthrough_ase" {
    convention      = "passthrough"
    name            = "TEST-DEV-ASE-RG"
    resource_type   = "azurerm_app_service_environment"
}

# App Service Plan
resource "azurecaf_naming_convention" "passthrough_plan" {
    convention      = "passthrough"
    name            = "TEST-DEV-PLAN-RG"
    resource_type   = "azurerm_app_service_plan"
}

# Azure SQL DB Server
resource "azurecaf_naming_convention" "passthrough_sql" {
    convention      = "passthrough"
    name            = "TEST-DEV-SQL-RG"
    resource_type   = "azurerm_sql_server"
}

# Azure SQL DB
resource "azurecaf_naming_convention" "passthrough_sqldb" {
    convention      = "passthrough"
    name            = "TEST-DEV-SQLDB-RG"
    resource_type   = "azurerm_sql_database"
}
`