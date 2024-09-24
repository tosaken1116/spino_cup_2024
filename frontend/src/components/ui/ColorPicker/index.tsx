

import { FC, MouseEvent, TouchEvent, useEffect, useRef, useState } from 'react';

type ColorPickerProps = {
  onChangeColor: (color: string) => void;
};

type RGB = {
  r: number;
  g: number;
  b: number;
};

export const ColorPicker: FC<ColorPickerProps> = ({ onChangeColor }) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const [selectedColor, setSelectedColor] = useState<string>('hsl(0, 100%, 50%)');
  const [hue, setHue] = useState<number>(0);
  const [saturation, setSaturation] = useState<number>(100);
  const [isDragging, setIsDragging] = useState<boolean>(false);

  const LIGHTNESS = 60;


  const hslToRgb = (h: number, s: number, l: number): RGB => {
    s /= 100;
    l /= 100;

    const c = (1 - Math.abs(2 * l - 1)) * s;
    const hh = h / 60;
    const x = c * (1 - Math.abs((hh % 2) - 1));
    let r1 = 0,
      g1 = 0,
      b1 = 0;

    if (0 <= hh && hh < 1) {
      r1 = c;
      g1 = x;
      b1 = 0;
    } else if (1 <= hh && hh < 2) {
      r1 = x;
      g1 = c;
      b1 = 0;
    } else if (2 <= hh && hh < 3) {
      r1 = 0;
      g1 = c;
      b1 = x;
    } else if (3 <= hh && hh < 4) {
      r1 = 0;
      g1 = x;
      b1 = c;
    } else if (4 <= hh && hh < 5) {
      r1 = x;
      g1 = 0;
      b1 = c;
    } else if (5 <= hh && hh < 6) {
      r1 = c;
      g1 = 0;
      b1 = x;
    }

    const m = l - c / 2;
    return {
      r: Math.round((r1 + m) * 255),
      g: Math.round((g1 + m) * 255),
      b: Math.round((b1 + m) * 255),
    };
  };


  const drawColorWheel = () => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    const width = canvas.width;
    const height = canvas.height;
    const radius = Math.min(width, height) / 2;
    const imageData = ctx.createImageData(width, height);
    const data = imageData.data;

    for (let y = 0; y < height; y++) {
      for (let x = 0; x < width; x++) {
        const dx = x - width / 2;
        const dy = y - height / 2;
        const distance = Math.sqrt(dx * dx + dy * dy);
        if (distance > radius) {

          data[(y * width + x) * 4 + 3] = 0;
          continue;
        }


        let currentHue = Math.atan2(dy, dx) * (180 / Math.PI);
        if (currentHue < 0) currentHue += 360;


        let currentSaturation = (distance / radius) * 100;
        currentSaturation = Math.min(Math.max(currentSaturation, 0), 100);


        const { r, g, b } = hslToRgb(currentHue, currentSaturation, LIGHTNESS);

        const index = (y * width + x) * 4;
        data[index] = r;
        data[index + 1] = g;
        data[index + 2] = b;
        data[index + 3] = 255;
      }
    }

    ctx.putImageData(imageData, 0, 0);
  };


  const handleSelection = (clientX: number, clientY: number) => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const rect = canvas.getBoundingClientRect();
    const x = clientX - rect.left;
    const y = clientY - rect.top;

    const width = canvas.width;
    const height = canvas.height;
    const dx = x - width / 2;
    const dy = y - height / 2;
    const distance = Math.sqrt(dx * dx + dy * dy);
    const radius = Math.min(width, height) / 2;

    if (distance > radius) {

      return;
    }


    let currentHue = Math.atan2(dy, dx) * (180 / Math.PI);
    if (currentHue < 0) currentHue += 360;


    let currentSaturation = (distance / radius) * 100;
    currentSaturation = Math.min(Math.max(currentSaturation, 0), 100);

    setHue(currentHue);
    setSaturation(currentSaturation);
    const newColor = `hsl(${currentHue.toFixed(0)}, ${currentSaturation.toFixed(0)}%, ${LIGHTNESS}%)`;
    setSelectedColor(newColor);
    onChangeColor(newColor);
  };


  const handleMouseDown = (e: MouseEvent<HTMLCanvasElement>) => {
    setIsDragging(true);
    handleSelection(e.clientX, e.clientY);
  };

  const handleMouseMove = (e: MouseEvent<HTMLCanvasElement>) => {
    if (isDragging) {
      handleSelection(e.clientX, e.clientY);
    }
  };

  const handleMouseUp = () => {
    setIsDragging(false);
  };

  const handleTouchStart = (e: TouchEvent<HTMLCanvasElement>) => {
    setIsDragging(true);
    const touch = e.touches[0];
    handleSelection(touch.clientX, touch.clientY);
  };

  const handleTouchMove = (e: TouchEvent<HTMLCanvasElement>) => {
    if (isDragging) {
      const touch = e.touches[0];
      handleSelection(touch.clientX, touch.clientY);
    }
  };

  const handleTouchEnd = () => {
    setIsDragging(false);
  };


  useEffect(() => {
    const handleGlobalMouseMove = (e: MouseEvent) => {
      if (isDragging) {
        handleSelection(e.clientX, e.clientY);
      }
    };

    const handleGlobalMouseUp = () => {
      setIsDragging(false);
    };

    const handleGlobalTouchMove = (e: TouchEvent) => {
      if (isDragging) {
        const touch = e.touches[0];
        handleSelection(touch.clientX, touch.clientY);
      }
    };

    const handleGlobalTouchEnd = () => {
      setIsDragging(false);
    };

    window.addEventListener('mousemove', handleGlobalMouseMove);
    window.addEventListener('mouseup', handleGlobalMouseUp);
    window.addEventListener('touchmove', handleGlobalTouchMove);
    window.addEventListener('touchend', handleGlobalTouchEnd);

    return () => {
      window.removeEventListener('mousemove', handleGlobalMouseMove);
      window.removeEventListener('mouseup', handleGlobalMouseUp);
      window.removeEventListener('touchmove', handleGlobalTouchMove);
      window.removeEventListener('touchend', handleGlobalTouchEnd);
    };
  }, [isDragging]);


  useEffect(() => {
    drawColorWheel();

  }, []);


  const drawSelectionMarker = () => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    const width = canvas.width;
    const height = canvas.height;
    const radius = Math.min(width, height) / 2;


    drawColorWheel();


    const angleRad = (hue * Math.PI) / 180;
    const markerX = width / 2 + Math.cos(angleRad) * (saturation / 100) * radius;
    const markerY = height / 2 + Math.sin(angleRad) * (saturation / 100) * radius;


    ctx.beginPath();
    ctx.arc(markerX, markerY, isDragging ? 15 : 10, 0, 2 * Math.PI);
    ctx.strokeStyle = '#000';
    ctx.lineWidth = 1;
    ctx.fillStyle = '#fff';
    ctx.fill();
    ctx.stroke();
  };


  useEffect(() => {
    drawSelectionMarker();

  }, [hue, saturation, isDragging]);

  return (
    <div className="relative flex flex-col items-center justify-center min-h-screen bg-gray-100 p-4">
      {/* ドラッグ中のみ選択色の大きな表示 */}
      {isDragging && (
        <div className="absolute top-10">
          <div
            className="w-32 h-32 rounded-full border-4"
            style={{ backgroundColor: selectedColor, borderColor: '#ccc' }}
          />
        </div>
      )}

      {/* カラーホイール */}
      <canvas
        ref={canvasRef}
        width={300}
        height={300}
        className="w-72 h-72 bg-white rounded-full shadow-md cursor-pointer"
        onMouseDown={handleMouseDown}
        onMouseMove={handleMouseMove}
        onMouseUp={handleMouseUp}
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onTouchEnd={handleTouchEnd}
        aria-label="HSL カラーピッカー"
        role="slider"
        tabIndex={0}
      />
    </div>
  );
};
