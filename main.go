package main

import (
	"3d-engine/shaders"
	"3d-engine/textures"
	"log"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	g_width  = 800
	g_height = 600
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("Could not init glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(g_width, g_height, "3D-Engine", nil, nil)
	if err != nil {
		log.Fatalln("Could not create a window:", err)
	}
	defer window.Destroy()

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	gl.Viewport(0, 0, int32(g_width), int32(g_height))

	window.SetFramebufferSizeCallback(framebuffer_size_callback)

	shader, err := shaders.CreateShaderProgram()
	if err != nil {
		log.Fatalln("Could not create a shader program:", err)
	}
	defer shader.Delete()

	// Test Rectangle
	vertices := []float32{
		// position // color // texture coords
		0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, // top right
		0.5, -0.5, 0.0, 0.0, 1.0, 0.0, 1.0, 0.0, // bottom right
		-0.5, -0.5, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, // bottom left
		-0.5, 0.5, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, // top left
	}

	indices := []uint32{
		0, 1, 3,
		1, 2, 3,
	}

	var vao, vbo, ebo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)
	defer gl.DeleteVertexArrays(1, &vao)
	defer gl.DeleteBuffers(1, &vbo)
	defer gl.DeleteBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	// vbo
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(sizeof[float32]()), gl.Ptr(vertices), gl.STATIC_DRAW)

	// ebo
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*int(sizeof[uint32]()), gl.Ptr(indices), gl.STATIC_DRAW)

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*int32(sizeof[float32]()), nil)
	gl.EnableVertexAttribArray(0)

	// color
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*int32(sizeof[float32]()), gl.Ptr(3*sizeof[float32]()))
	gl.EnableVertexAttribArray(1)

	// texture
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*int32(sizeof[float32]()), gl.Ptr(6*sizeof[float32]()))
	gl.EnableVertexAttribArray(2)

	// unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
	// End Test Rectangle

	// Texture
	var textureId uint32
	gl.GenTextures(1, &textureId)
	gl.BindTexture(gl.TEXTURE_2D, textureId)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	texture := textures.Load("wall")

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, texture.Width, texture.Height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(texture.Data))
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.BindTexture(gl.TEXTURE_2D, 0)
	// End Texture

	// Wireframe
	// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)

	for !window.ShouldClose() {
		processInput(window)

		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.BindTexture(gl.TEXTURE_2D, textureId)

		shader.Use()

		gl.BindVertexArray(vao)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
		gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, nil)
		// gl.DrawArrays(gl.TRIANGLES, 0, 3)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
}

func framebuffer_size_callback(window *glfw.Window, width, height int) {
	g_width = width
	g_height = height

	gl.Viewport(0, 0, int32(width), int32(height))
}

func sizeof[T any]() uintptr {
	var v T
	return unsafe.Sizeof(v)
}
