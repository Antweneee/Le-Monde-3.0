import { extendTheme } from '@chakra-ui/react';

import Text from './components/text';
import Button from './components/button';
import Input from './components/input';
import breakpoints from './foundations/breakpoints';
import colors from './foundations/colors';
import fontWeights from './foundations/fontWeights';
import borderRadius from './foundations/borderRadius';

const overrides = {
	colors,
	fontWeights,
	...borderRadius,
	breakpoints,
	components: {
		Text,
		Button,
		Input,
	},
};

export default extendTheme(overrides, breakpoints);
