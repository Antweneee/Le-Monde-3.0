import * as React from 'react';
import { Box, HStack, Icon, Text, VStack } from '@chakra-ui/react';
import { ChevronLeftIcon, ChevronRightIcon } from '@chakra-ui/icons';
import { FaBookOpen, FaPenNib } from 'react-icons/fa';
import { MdAccountCircle } from 'react-icons/md';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const Title = ({ icon, name }: { icon: any; name: string }): JSX.Element => (
	<Box position="relative" w="100%">
		<Icon as={icon} position="absolute" left="40px" boxSize={12} color="black" />
		<Text variant="h4" pl="120px">
			{name}
		</Text>
	</Box>
);

const Option = ({ name, isSelected }: { name: string; isSelected: boolean }): JSX.Element => (
	<Box position="relative" w="100%" cursor="pointer">
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

const Dashboard = (): JSX.Element => (
	<>
		<VStack
			bg="primary.4"
			w="360px"
			h="100vh"
			position="sticky"
			spacing="48px"
			p="16px 16px 16px 0px"
			borderTopRightRadius="sm"
			borderBottomRightRadius="sm"
		>
			<HStack w="100%" justify="flex-end">
				<ChevronLeftIcon boxSize={6} color="black" />
				<ChevronRightIcon boxSize={6} color="black" />
			</HStack>
			<Text variant="h4">@username</Text>
			<VStack align="start" w="100%">
				<Title icon={FaBookOpen} name="Lire" />
				<Option name="Favoris" isSelected={true} />
				<Option name="Marque-pages" isSelected={false} />
				<Option name="Nouveautés" isSelected={false} />
				<Option name="Découvertes" isSelected={false} />
			</VStack>
			<VStack align="start" w="100%">
				<Title icon={FaPenNib} name="Écrire" />
				<Option name="Nouvel article" isSelected={false} />
				<Option name="Publications" isSelected={false} />
				<Option name="Brouillons" isSelected={false} />
				<Option name="Statistiques" isSelected={false} />
			</VStack>
			<VStack align="start" w="100%">
				<Title icon={MdAccountCircle} name="Compte" />
				<Option name="Réglages" isSelected={false} />
				<Option name="Déconnexion" isSelected={false} />
			</VStack>
		</VStack>
	</>
);

export default Dashboard;
