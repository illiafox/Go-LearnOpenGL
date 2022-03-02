package main

import "C"
import (
	"github.com/go-gl/gl/v3.3-compatibility/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
	"os"
	"runtime"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()

	// We need this in Go
	err := gl.Init()
	if err != nil {
		panic(err)
	}
}

// settings
const SCR_WIDTH int = 800
const SCR_HEIGHT int = 600

const vertexShaderSource = `#version 330 core
layout (location = 0) in vec3 aPos;
void main()
{
   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
}`

const fragmentShaderSource = `#version 330 core
out vec4 FragColor;
void main()
{
   FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
}`

func main() {
	// glfw: initialize and configure
	// ------------------------------
	glfw.Init()
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	if runtime.GOOS == "darwin" {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}

	// glfw window creation
	// --------------------
	window, err := glfw.CreateWindow(SCR_WIDTH, SCR_HEIGHT, "LearnOpenGL", nil, nil)
	if err != nil {
		log.Fatalln("Failed to create GLFW window\n", err)
		glfw.Terminate()
		os.Exit(-1)
	}
	window.MakeContextCurrent()
	window.SetFramebufferSizeCallback(framebuffer_size_callback)

	// build and compile our shader program
	// ------------------------------------
	// vertex shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, 1, C.CString(vertexShaderSource), nil)
	// render loop
	// -----------
	for !window.ShouldClose() {
		// input
		// -----
		processInput(window)

		// glfw: swap buffers and poll IO events (keys pressed/released, mouse moved etc.)
		// -------------------------------------------------------------------------------
		window.SwapBuffers()
		glfw.PollEvents()
	}

	// glfw: terminate, clearing all previously allocated GLFW resources.
	// ------------------------------------------------------------------
	glfw.Terminate()
	return
}

// process all input: query GLFW whether relevant keys are pressed/released this frame and react accordingly
// ---------------------------------------------------------------------------------------------------------
func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

// glfw: whenever the window size changed (by OS or user resize) this callback function executes
// ---------------------------------------------------------------------------------------------
func framebuffer_size_callback(window *glfw.Window, width int, heigh int) {
	// make sure the viewport matches the new window dimensions; note that width and
	// height will be significantly larger than specified on retina displays.
	gl.Viewport(0, 0, int32(width), int32(heigh))
}
