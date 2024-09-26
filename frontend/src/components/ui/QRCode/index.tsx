import { QRCodeCanvas } from "qrcode.react";
import type { FC } from "react";

type Props = {
	url: string;
};

export const QRCode: FC<Props> = (props) => {
	return (
		<QRCodeCanvas
			value={props.url}
			size={256}
			bgColor={"#000000"}
			fgColor={"#ffffff"}
			level={"L"}
			includeMargin={false}
			imageSettings={{
				src: "/favicon.ico",
				x: undefined,
				y: undefined,
				height: 24,
				width: 24,
				excavate: true,
			}}
		/>
	);
};

export default QRCode;
