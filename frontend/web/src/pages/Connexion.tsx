import * as React from 'react';
import { Button, Input, Link } from '@chakra-ui/react';
import { Link as RouteLink, useNavigate } from 'react-router-dom';

const Connexion = (): JSX.Element => {
	const navigate = useNavigate();

	return (
		<>
			<Input id="connexion-email-input" variant="primary-1" placeholder="e-mail ou nom d'utilisateur" />
			<Input id="connexion-pwd-input" variant="primary-1" placeholder="mot de passe" type="password" />
			<Button id="connexion-connexion-btn" variant="primary-1" onClick={() => navigate('/favoris')}>
				Connexion
			</Button>
			<Link as={RouteLink} to="/inscription" w="100%">
				<Button id="connexion-inscription-btn" variant="secondary-4">
					Inscription
				</Button>
			</Link>
		</>
	);
};

export default Connexion;
