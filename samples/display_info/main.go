package main

import (
	"fmt"
	"github.com/gonutz/d3d9"
)

func main() {
	check(d3d9.Init())
	defer d3d9.Close()

	d3d, err := d3d9.Create(d3d9.SDK_VERSION)
	check(err)
	defer d3d.Release()

	adapterCount := d3d.GetAdapterCount()
	fmt.Println(adapterCount, "adapter(s) detected")

	for adapter := uint(0); adapter < adapterCount; adapter++ {
		modeCount := d3d.GetAdapterModeCount(adapter, d3d9.FMT_X8R8G8B8)
		fmt.Printf("adapter %v has %v R8G8B8 modes:\n", adapter, modeCount)

		for mode := uint(0); mode < modeCount; mode++ {
			adapaterMode, err := d3d.EnumAdapterModes(
				adapter,
				d3d9.FMT_X8R8G8B8,
				mode,
			)
			check(err)
			fmt.Printf(
				"  mode %2d: %vx%v %vHz\n",
				mode,
				adapaterMode.Width,
				adapaterMode.Height,
				adapaterMode.RefreshRate,
			)
		}

		displayMode, err := d3d.GetAdapterDisplayMode(adapter)
		check(err)
		fmt.Printf(
			"current display mode: %vx%v %vHz, format %v\n",
			displayMode.Width,
			displayMode.Height,
			displayMode.RefreshRate,
			displayMode.Format,
		)

		id, err := d3d.GetAdapterIdentifier(adapter, 0)
		check(err)
		fmt.Printf(
			"driver: %v (%v)\n",
			string(id.Driver[:]),
			string(id.Description[:]),
		)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
