package web

import (
  "html/template"
  "io/ioutil"
  "path"
)

var TEMPLATE_FOLDER = "templates"
var TEMPLATE_LEFT_DELIMITER = "{!{"
var TEMPLATE_RIGHT_DELIMITER = "}!}"

var templates *template.Template


type templateRecord struct {
  name string
  file string
}
var templateRecords []templateRecord



func init() {
  templates = template.New("__unused__")
}



func registerTemplate(fname, templateName string)error {
  fname = path.Join(TEMPLATE_FOLDER, fname)
  templateRecords = append(templateRecords, templateRecord{name: templateName, file: fname,})

  templ := templates.New(templateName)
  templ.Delims(TEMPLATE_LEFT_DELIMITER, TEMPLATE_RIGHT_DELIMITER)

  fdata, err := ioutil.ReadFile(fname)
  if err != nil{
    return err
  }
  _, err = templ.Parse(string(fdata))
  if err != nil{
    return err
  }

  return nil
}
