import clsx from "clsx";
import { useState } from "react";
import { SquareArrowDownRight } from "../../../../components/icons/SquareArrowDownRight";
import { SquareArrowUpLeft } from "../../../../components/icons/SquareArrowUpLeft";
import { ChangePenSizeSlider } from "../../../../components/ui/ChangePenSizeSlider";
import { ColorPicker } from "../../../../components/ui/ColorPicker";
import { DrawButton } from "../../../../components/ui/DrawButton";
import { useOrientationCalculate } from "../../../../libs/orientationCalculate";
import type { UserAction } from "../../../../libs/wsClients";

type Props = Omit<UserAction, "type">;

export const UserController = ({
  handleChangePointerColor,
  screenSize,
  handleClickPointer,
  handleChangePenSize,
  handleChangePointerPosition,
}: Props) => {
  const [position, setPosition] = useState({
    x: 0,
    y: 0,
  });
  const {
    handlePermissionGranted,
    handleSetLeftTopPoint,
    handleSetRightBottomPoint,
    permissionGranted,
  } = useOrientationCalculate({
    screenSize,
    handleChangePointerPosition: (props) => {
      handleChangePointerPosition(props);
      setPosition(props);
    },
  });
  return (
    <div className="flex flex-col gap-4 h-full">
      {!permissionGranted && (
        <div className="w-full absolute h-screen bg-slate-500/70 z-50">
          <div className="flex flex-col items-center absolute  left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2  justify-center gap-2 border-2 rounded-xl p-4 w-fit">
            <p>情報を取得する権限が必要です</p>

            <button
              type="button"
              className="w-fit rounded-xl border-2 bg-green-300 text-slate-600 font-semibold ml-4 py-2"
              onClick={handlePermissionGranted}
            >
              オンにする
            </button>
          </div>
        </div>
      )}
      <div className="flex flex-row gap-2 h-full">
        <div className="w-full flex justify-center items-center">
          <div className="relative w-48 h-28 rounded-md border-2">
            <div
              className="absolute w-2 h-2"
              style={{
                left: `${position.x}px`,
                top: `${position.y}px`,
              }}
            />
          </div>
        </div>
        <div className="w-1/4">
          <ChangePenSizeSlider onChange={handleChangePenSize} />
        </div>
      </div>
      <div className="flex flex-row items-end h-full p-4">
        <div className="w-2/3">
          <ColorPicker onChangeColor={handleChangePointerColor} />
        </div>
        <div className="w-1/3 flex flex-col gap-4 items-center">
          <InitializeButton onClick={handleSetLeftTopPoint}>
            <SquareArrowUpLeft className="w-8 h-8" />
          </InitializeButton>
          <InitializeButton onClick={handleSetRightBottomPoint}>
            <SquareArrowDownRight className="w-8 h-8" />
          </InitializeButton>
          <DrawButton onChangePointer={handleClickPointer} />
        </div>
      </div>
    </div>
  );
};

const InitializeButton = ({
  className,
  children,
  ...props
}: React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
>) => {
  return (
    <button
      type="button"
      className={clsx(
        "shadow-md shadow-black rounded-full w-16 h-16 text-blue-800 border-4 border-blue-400 flex justify-center items-center",
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};
