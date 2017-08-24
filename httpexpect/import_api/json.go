package import_api

var validJSON string = `
{
  "recipe": "hello",
  "state": "new",
  "files": [
	{
	  "alias_name": "v4",
	  "url": "https://s3-eu-west-1.amazonaws.com/dp-publish-content-test/OCIGrowth.csv"
	}
  ]
}`

var invalidJSON string = `
{
  "number_of_instances": "2",
  "state": "New",
  "files": [
	{
	  "alias_name": "v4",
	  "url": "https://s3-eu-west-1.amazonaws.com/dp-publish-content-test/OCIGrowth.csv"
	}
  ]
}`

var validJobStateJSON string = `
{
  "state": "created"
}`

var invalidJobStateJSON string = `
{
  "not_the_state_you_were_expecting": "broken",
}`

var validFileJSON string = `
{
  "alias_name": "v4",
  "url": "https://s3-eu-west-1.amazonaws.com/dp-publish-content-test/OCIGrowth.csv"
}`

var invalidFileJSON string = `
{
  "url": "https://s3-eu-west-1.amazonaws.com/dp-publish-content-test/OCIGrowth.csv"
}`
