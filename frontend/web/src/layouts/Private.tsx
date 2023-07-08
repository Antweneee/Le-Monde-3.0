import * as React from 'react';
import { Box, HStack, Icon, Text, VStack } from '@chakra-ui/react';
import { ChevronLeftIcon, ChevronRightIcon } from '@chakra-ui/icons';
import { FaBookOpen, FaPenNib } from 'react-icons/fa';
import { MdAccountCircle } from 'react-icons/md';
import { useLocation, useNavigate } from 'react-router-dom';

type PrivateProps = { children: JSX.Element };

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const Title = ({ icon, name }: { icon: any; name: string }): JSX.Element => (
	<Box position="relative" w="100%">
		<Icon as={icon} position="absolute" left="40px" boxSize={12} color="black" />
		<Text variant="h4" pl="120px">
			{name}
		</Text>
	</Box>
);

const Option = ({
	name,
	isSelected,
	onClick,
}: {
	name: string;
	isSelected: boolean;
	onClick: () => void;
}): JSX.Element => (
	<Box position="relative" w="100%" cursor="pointer" onClick={onClick}>
		<Box
			position="absolute"
			left={0}
			display={isSelected ? 'block' : 'none'}
			w="8px"
			h="24px"
			bg="primary.1"
			borderRadius="sm"
		/>
		<Text variant="link" pl="120px" color={isSelected ? 'primary.1' : 'black'}>
			{name}
		</Text>
	</Box>
);

const NavBar = (): JSX.Element => {
	const location = useLocation();
	const navigate = useNavigate();

	return (
		<VStack
			bg="primary.4"
			w="360px"
			minW="360px !important"
			h="100vh"
			position="sticky"
			spacing="48px"
			p="16px 16px 16px 0px"
			borderTopRightRadius="sm"
			borderBottomRightRadius="sm"
		>
			<HStack w="100%" justify="flex-end">
				<ChevronLeftIcon boxSize={6} color="black" cursor="grab" onClick={() => navigate(-1)} />
				<ChevronRightIcon boxSize={6} color="black" cursor="grab" onClick={() => navigate(1)} />
			</HStack>
			<Text variant="h4">@username</Text>
			<VStack align="start" w="100%">
				<Title icon={FaBookOpen} name="Lire" />
				<Option name="Favoris" isSelected={location.pathname === '/favoris'} onClick={() => navigate('/favoris')} />
				<Option
					name="Marque-pages"
					isSelected={location.pathname === '/marque-pages'}
					onClick={() => navigate('/marque-pages')}
				/>
				<Option
					name="Nouveautés"
					isSelected={location.pathname === '/nouveautes'}
					onClick={() => navigate('/nouveautes')}
				/>
				<Option
					name="Découvertes"
					isSelected={location.pathname === '/decouvertes'}
					onClick={() => navigate('/decouvertes')}
				/>
			</VStack>
			<VStack align="start" w="100%">
				<Title icon={FaPenNib} name="Écrire" />
				<Option
					name="Nouvel article"
					isSelected={location.pathname === '/nouvel-article'}
					onClick={() => navigate('/nouvel-article')}
				/>
				<Option
					name="Publications"
					isSelected={location.pathname === '/publications'}
					onClick={() => navigate('/publications')}
				/>
				<Option
					name="Brouillons"
					isSelected={location.pathname === '/brouillons'}
					onClick={() => navigate('/brouillons')}
				/>
				<Option
					name="Statistiques"
					isSelected={location.pathname === '/statistiques'}
					onClick={() => navigate('/statistiques')}
				/>
			</VStack>
			<VStack align="start" w="100%">
				<Title icon={MdAccountCircle} name="Compte" />
				<Option name="Réglages" isSelected={location.pathname === '/reglages'} onClick={() => navigate('/reglages')} />
				<Option
					name="Déconnexion"
					isSelected={location.pathname === '/deconnexion'}
					onClick={() => navigate('/deconnexion')}
				/>
			</VStack>
		</VStack>
	);
};

const Private = ({ children }: PrivateProps): JSX.Element => (
	<HStack align="start">
		<NavBar />
		{children}
	</HStack>
);

export default Private;
