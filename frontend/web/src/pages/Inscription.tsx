import * as React from 'react';
import { Button, Input, Link } from '@chakra-ui/react';
import { Link as RouteLink, useNavigate } from 'react-router-dom';

const Inscription = (): JSX.Element => {
	const navigate = useNavigate();

	return (
		<>
			<Input variant="primary-1" placeholder="e-mail" />
			<Input variant="primary-1" placeholder="nom d'utilisateur" />
			<Input variant="primary-1" placeholder="mot de passe" type="password" />
			<Input variant="primary-1" placeholder="confirmation du mot de passe" type="password" />
			<Button variant="primary-1" onClick={() => navigate('/favoris')}>
				Inscription
			</Button>
			<Link as={RouteLink} to="/connexion" w="100%">
				<Button variant="secondary-4">Connexion</Button>
			</Link>
		</>
	);
};

export default Inscription;
