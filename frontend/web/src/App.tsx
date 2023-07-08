import * as React from 'react';
import { ChakraProvider } from '@chakra-ui/react';

import Routes from 'routes/Index';

import theme from 'theme';
import 'theme/index.css';

export const App = (): JSX.Element => (
	<ChakraProvider theme={theme}>
		<Routes />
	</ChakraProvider>
);
