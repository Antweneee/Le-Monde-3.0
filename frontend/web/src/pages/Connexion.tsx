import * as React from 'react';
import { Button, Input, Link } from '@chakra-ui/react';
import { Link as RouteLink } from 'react-router-dom';

const Connexion = (): JSX.Element => (
	<>
		<Input variant="primary-1" placeholder="e-mail ou nom d'utilisateur" />
		<Input variant="primary-1" placeholder="mot de passe" />
		<Button variant="primary-1">Connexion</Button>
		<Link as={RouteLink} to="/inscription" w="100%">
			<Button variant="secondary-4">Inscription</Button>
		</Link>
	</>
);

export default Connexion;
