import { useCallback, useEffect, useState } from "react";
import type { ScreenSize } from "../../generated/wsClient/room/model";
import { calculateScreenPosition } from "../calculateScreenPosition";

export const useOrientationCalculate = ({
  handleChangePointerPosition,
  screenSize,
}: {
  screenSize: ScreenSize;
  handleChangePointerPosition: (props: { x: number; y: number }) => void;
}) => {
  const [leftTopOrientation, setLeftTopOrientation] = useState({
    alpha: 0,
    beta: 0,
  });
  const [permissionGranted, setPermissionGranted] = useState(false);
  const [rightBottomOrientation, setRightTopOrientation] = useState({
    alpha: 0,
    beta: 0,
  });

  const handleSetLeftTopPoint = useCallback(() => {
    window.addEventListener(
      "deviceorientation",
      (e) => {
        setLeftTopOrientation({
          alpha: e.alpha ?? 0,
          beta: e.beta ?? 0,
        });
      },
      {
        once: true,
      }
    );
  }, []);

  const handleSetRightBottomPoint = useCallback(() => {
    window.addEventListener(
      "deviceorientation",
      (e) => {
        setRightTopOrientation({
          alpha: e.alpha ?? 0,
          beta: e.beta ?? 0,
        });
      },
      {
        once: true,
      }
    );
  }, []);

  const setCurrentPointer = useCallback(
    (props: { alpha: number; beta: number }) => {
      handleChangePointerPosition(
        calculateScreenPosition({
          current: props,
          max: leftTopOrientation,
          min: rightBottomOrientation,
          screenSize,
        })
      );
    },
    [
      handleChangePointerPosition,
      leftTopOrientation,
      rightBottomOrientation,
      screenSize,
    ]
  );

  const handlePermissionGranted = async () => {
    // @ts-ignore
    await DeviceOrientationEvent.requestPermission()
      // @ts-ignore
      .then((permissionState) => {
        if (permissionState === "granted") {
          setPermissionGranted(true);
          // @ts-ignore
          window.addEventListener("deviceorientation", setCurrentPointer);
        }
      })
      .catch(console.error);
  };

  useEffect(() => {
    // @ts-ignore
    if (typeof DeviceMotionEvent.requestPermission === "function") {
      // @ts-ignore
      DeviceMotionEvent.requestPermission()
        // @ts-ignore
        .then((permissionState) => {
          if (permissionState === "granted") {
            setPermissionGranted(true);
            // @ts-ignore
            window.addEventListener("deviceorientation", setCurrentPointer);
          }
        })
        .catch(console.error);
    } else {
      setPermissionGranted(true);
      // @ts-ignore
      window.addEventListener("deviceorientation", setCurrentPointer);
    }

    return () => {
      // @ts-ignore
      window.removeEventListener("deviceorientation", setCurrentPointer);
    };
  }, [setCurrentPointer]);
  return {
    leftTopOrientation,
    rightBottomOrientation,
    permissionGranted,
    handleSetLeftTopPoint,
    handleSetRightBottomPoint,
    handlePermissionGranted,
  };
};
