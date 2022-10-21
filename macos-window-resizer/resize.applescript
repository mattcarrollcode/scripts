-- This script was written to put Chrome and VS Code in a specific location with a specific
-- window size in order to better setup screen recording with OBS for livestreaming on macOS.
-- Setup:
--   * macOS with 4k screen
--   * scaling set to "looks like 2560 x 1440"
--   * OBS set to screen Displa Capture
--     * Crop: Manual
--     * Crop left: 0
--     * Crop top: 900
--     * Crop right: 3200
--     * Crop bottom: 900
-- For example see example.png and exmaple2.png in this directory.
-- Extensive tweaking to the appHeight, appWidth, xAxis, yAxis, xAxisOffset, and yAxis offset
-- is required for different applications, monitor resolutions, scaling, and OBS settings.
-- To run this script type `osascript resize.applescript` and press enter.

-- Find resolution of monitor
tell application "Finder"
	set screenResolution to bounds of window of desktop
end tell
set screenWidth to item 3 of screenResolution
set screenHeight to item 4 of screenResolution

-- Set Chrome to be 1080p at the middle of the left side of the monitor
tell application "Google Chrome"
	set appHeight to 600 -- 540 + 60 offset to remove top window chrome and bottom scroll bar
	set appWidth to 975 -- 960 + 15 offset to remove right scroll bar
	set yAxisOffset to -10 -- offset yAxis by -10 to account for top window chrome and bottom scroll bar
	set yAxis to ((screenHeight - appHeight) / 2 as integer) + yAxisOffset -- put the app in the middle of the screen vertically
	set xAxisOffset to 0
	set xAxis to 0 + xAxisOffset -- put app to the right hand side of the screen
	activate -- Bring the app to the foreground
	set the bounds of the first window to {xAxis, yAxis, appWidth + xAxis, appHeight + yAxis}
end tell

-- Set VS Code to be 1080p at the middle of the left side of the monitor
tell application "Visual Studio Code" to activate -- Bring the app to the foreground
tell application "System Events"
    tell (the first process where it is frontmost) -- control the frontmost app
	    set appHeight to 580 -- 540 + 40 offset to remove top window chrome
		set appWidth to 960 -- 960 "px" = 1080 acutal pixels
		set yAxisOffset to -20 -- offset yAxis by -10 to account for top window chrome and bottom scroll bar
		set yAxis to ((screenHeight - appHeight) / 2 as integer) + yAxisOffset -- put the app in the middle of the screen vertically
		set xAxisOffset to 0
		set xAxis to 0 + xAxisOffset -- put app to the right hand side of the screen
        tell the front window
            set its size to [appWidth, appHeight]
            set its position to [xAxis,yAxis]
        end tell
    end tell
end tell