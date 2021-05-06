import 'package:flutter/material.dart';
import 'package:walking_octopus/components/icon_btn_with_counter.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/size_config.dart';

import 'home_body.dart';

class HomeScreen extends StatelessWidget {
  static String routeName = "/home";
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Body(),
    );
  }
}