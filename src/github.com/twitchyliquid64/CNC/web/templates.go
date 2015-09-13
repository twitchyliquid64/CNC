package web

import (
  "github.com/twitchyliquid64/CNC/logging"
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


func templateReInit() {
  logging.Info("web", "Now reloading all templates.")
  templates = template.New("__unused__")
  for _, tempFile := range templateRecords {
    logging.Info("web", "Loading template: ", tempFile.name)
    if err :=newTemplateFromFile(tempFile.file, tempFile.name); err != nil {
      logging.Error("web", "Template error: ", err)
    }
  }
}


func registerTemplate(fname, templateName string)error {
  fname = path.Join(TEMPLATE_FOLDER, fname)
  templateRecords = append(templateRecords, templateRecord{name: templateName, file: fname,})

  return newTemplateFromFile(fname, templateName)
}

func newTemplateFromFile(fname, templateName string)error {
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
