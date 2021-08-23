package main

import {
  "fnt"
  // "github.com/kagani/go-face"
}

 const datadir = "images"

func main() {
  fnt.Println("Facial Recognition System V0.0.1")

  rec, err := face.NewRecognizer(datadir)

  if err != nil {
    fnt.Println("Error in creting NewRecognizer")
    fnt.Println(err)
  }
  defer rec.Close()

  IronMan := filepath.Join(dataDir, "IronMan.jpg")
  // Recognize faces on that image.
  faces, err := rec.RecognizeFile(testImagePristin)
  if err != nil {
    log.Fatalf("Can't recognize: %v", err)
  }
  if len(faces) != 10 {
    log.Fatalf("Wrong number of faces")
  }
  var samples []face.Descriptor
  var cats []int32
  for i, f := range faces {
    samples = append(samples, f.Descriptor)
    // Each face is unique on that image so goes to its own
    // category.
    cats = append(cats, int32(i))
  }
  // Name the categories, i.e. people on the image.
  labels := []string{
    "IronMan", "Thor", "Captain America", "Vulture", "Nick",
    "Flash", "Thanos", "Hulk", "Spider Man", "Natasha Romanoff",
  }
  // Pass samples to the recognizer.
  rec.SetSamples(samples, cats)
  // Now let's try to classify some not yet known image.
  testImageNayoung := filepath.Join(dataDir, "nayoung.jpg")
  nayoungFace, err := rec.RecognizeSingleFile(testImageNayoung)
  if err != nil {
    log.Fatalf("Can't recognize: %v", err)
  }
  if nayoungFace == nil {
    log.Fatalf("Not a single face on the image")
  }
  catID := rec.Classify(nayoungFace.Descriptor)
  if catID < 0 {
    log.Fatalf("Can't classify")
  }
  // Finally print the classified label. It should be "Nayoung".
  fmt.Println(labels[catID])
}