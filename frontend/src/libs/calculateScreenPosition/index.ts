type Props = {
  max: {
    alpha: number;
    beta: number;
  };
  min: {
    alpha: number;
    beta: number;
  };
  center?: {
    alpha: number;
    beta: number;
  };
  current: {
    alpha: number;
    beta: number;
  };
  screenSize: {
    width: number;
    height: number;
  };
};

export const calculateScreenPosition = (props: Props) => {
  const { max, min, current } = props;

  const x = Math.tan(((max.alpha - current.alpha) / 180) * Math.PI);

  const y = (max.beta - current.beta) / (max.beta - min.beta);
  return { x: Math.min(Math.max(x, 0), 1), y: Math.min(Math.max(y, 0), 1) };
};
