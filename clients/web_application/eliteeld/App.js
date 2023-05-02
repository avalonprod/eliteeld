import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View } from 'react-native';

import Header from "./components/header/header";


export default function App() {
  return (
    <View style={styles.container}>
       <StatusBar style={'light'} />
      <Header></Header>
      <View style={styles.main}>
        
      </View> 
     
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },

  main: {
    height: '100%',
    backgroundColor: "#1E1E1E"
  }
});
