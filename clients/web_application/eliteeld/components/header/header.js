import { View, StyleSheet } from "react-native"


const Header = () => {
    return (
        <View style={styles.header}>

        </View>
    )
}

const styles = StyleSheet.create({
    header: {
        width: "100%",
        height: 100,
        position: "absolute",
        zIndex: 100,
        backgroundColor: "rgba(0, 0, 0, 0.8)",
        backdropFilter: 'blur(20px)',
        borderWidth: 1,
        borderBottomColor: "rgba(255, 255, 255, 0.5)"
    }
});

export default Header;