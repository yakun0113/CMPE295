import 'package:flutter/material.dart';
import 'package:walking_octopus/routs.dart';
import 'package:walking_octopus/screens/launch_screen.dart';
import 'package:walking_octopus/theme.dart';


void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: buildThemeData(),
      //home: SplashScreen(),
      initialRoute: LaunchScreen.routeName,
      routes: routes,
    );
  }
}