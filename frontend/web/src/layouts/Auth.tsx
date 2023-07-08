import * as React from 'react';
import { Text, VStack } from '@chakra-ui/react';

type AuthProps = { children: JSX.Element };

const Auth = ({ children }: AuthProps): JSX.Element => (
	<VStack spacing="56px" mt="132px">
		<VStack spacing="0px">
			<Text id="app-title" variant="h1" textAlign="center">
				Le Monde 3.0
			</Text>
			<Text id="app-description" variant="h5" textAlign="center">
				Le journal décentralisé luttant contre la censure.
			</Text>
		</VStack>
		<VStack w="496px" pb="64px">
			{children}
		</VStack>
	</VStack>
);

export default Auth;
