import 'package:flutter/material.dart';
import 'package:walking_octopus/size_config.dart';
import 'launch_body.dart';

class LaunchScreen extends StatelessWidget {
  static String routeName = "/splash";
  @override
  Widget build(BuildContext context) {
    // You have to call it on your starting screen
    SizeConfig().init(context);
    return Scaffold(
      body: Body(),
    );
  }
}