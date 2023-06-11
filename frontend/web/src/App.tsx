import * as React from 'react';
import { ChakraProvider, Text, theme } from '@chakra-ui/react';

export const App = (): JSX.Element => (
	<ChakraProvider theme={theme}>
		<Text>Welcome to Le Monde 3.0</Text>
	</ChakraProvider>
);
