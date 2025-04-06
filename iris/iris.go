package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	ort "github.com/yalue/onnxruntime_go"
)

func getDefaultSharedLibPath() string {
	if runtime.GOOS == "windows" {
		if runtime.GOARCH == "amd64" {
			return "../third_party/onnxruntime.dll"
		}
	}
	if runtime.GOOS == "darwin" {
		if runtime.GOARCH == "arm64" {
			return "../third_party/onnxruntime_arm64.dylib"
		}
		if runtime.GOARCH == "amd64" {
			return "../third_party/onnxruntime_amd64.dylib"
		}
	}
	if runtime.GOOS == "linux" {
		if runtime.GOARCH == "arm64" {
			return "../third_party/onnxruntime_arm64.so"
		}
		return "../third_party/onnxruntime.so"
	}
	fmt.Printf("Unable to determine a path to the onnxruntime shared library"+
		" for OS \"%s\" and architecture \"%s\".\n", runtime.GOOS,
		runtime.GOARCH)
	return ""
}

func runTest(onnxruntimeLibPath string) error {
	ort.SetSharedLibraryPath(onnxruntimeLibPath)
	e := ort.InitializeEnvironment()
	if e != nil {
		return fmt.Errorf("error initializing the onnxruntime library: %w", e)
	}
	defer ort.DestroyEnvironment()

	// Update input tensor for the Random Forest model
	inputData := []float32{5.1, 3.5, 1.4, 0.2} // Example Iris dataset input
	inputTensor, e := ort.NewTensor(ort.NewShape(1, 4), inputData)
	if e != nil {
		return fmt.Errorf("error creating the input tensor: %w", e)
	}
	defer inputTensor.Destroy()

	// Update output tensor for the Random Forest model
	outputTensor, e := ort.NewEmptyTensor[float32](ort.NewShape(1, 3)) // 3 classes
	if e != nil {
		return fmt.Errorf("error creating the output tensor: %w", e)
	}
	defer outputTensor.Destroy()

	// Update the ONNX file path and tensor names
	session, e := ort.NewAdvancedSession("./iris.onnx",
		[]string{"input"},  // Replace with the actual input name in your ONNX model
		[]string{"output"}, // Replace with the actual output name in your ONNX model
		[]ort.ArbitraryTensor{inputTensor},
		[]ort.ArbitraryTensor{outputTensor},
		nil)
	if e != nil {
		return fmt.Errorf("error creating the session: %w", e)
	}
	defer session.Destroy()

	e = session.Run()
	if e != nil {
		return fmt.Errorf("error executing the network: %w", e)
	}

	outputData := outputTensor.GetData()
	fmt.Printf("The network ran without errors.\n")
	fmt.Printf("  Input data: %v\n", inputData)
	fmt.Printf("  Predicted class probabilities: %v\n", outputData)
	return nil
}

func run() int {
	var onnxruntimeLibPath string
	flag.StringVar(&onnxruntimeLibPath, "onnxruntime_lib",
		getDefaultSharedLibPath(),
		"The path to the onnxruntime shared library for your system.")
	flag.Parse()
	if onnxruntimeLibPath == "" {
		fmt.Println("You must specify a path to the onnxruntime shared " +
			"on your system. Run with -help for more information.")
		return 1
	}
	e := runTest(onnxruntimeLibPath)
	if e != nil {
		fmt.Printf("Encountered an error running the network: %s\n", e)
		return 1
	}
	fmt.Printf("The network seemed to run OK!\n")
	return 0
}

func main() {
	os.Exit(run())
}
