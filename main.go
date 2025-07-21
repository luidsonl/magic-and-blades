package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Tela 2D para 3D - raylib-go")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	inMenu := true

	camera := rl.Camera3D{
		Position:   rl.NewVector3(5, 5, 5),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45.0,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		if inMenu {
			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText("Clique para começar o jogo", 220, 280, 20, rl.DarkGray)
			rl.EndDrawing()

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				inMenu = false
			}

		} else {
			rl.UpdateCamera(&camera, rl.CameraFirstPerson)

			rl.BeginDrawing()
			rl.ClearBackground(rl.RayWhite)

			rl.BeginMode3D(camera)
			rl.DrawCube(rl.NewVector3(0, 0, 0), 2, 2, 2, rl.Red)
			rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 2, 2, 2, rl.Maroon)
			rl.DrawGrid(10, 1)
			rl.EndMode3D()

			rl.DrawText("Use WASD + mouse para mover a câmera", 10, 10, 20, rl.DarkGray)
			rl.EndDrawing()
		}
	}
}
