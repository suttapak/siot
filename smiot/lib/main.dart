// ignore_for_file: prefer_const_constructors

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/auth_bloc.dart';
import 'package:smiot/bloc/box_bloc.dart';
import 'package:smiot/bloc/control_bloc.dart';
import 'package:smiot/bloc/display_bloc.dart';
import 'package:smiot/bloc/display_data_bloc.dart';
import 'package:smiot/bloc/user_bloc.dart';
import 'package:smiot/screens/home_screen.dart';
import 'package:smiot/utils/auth_guard.dart';

void main() => runApp(MainApp());

class MainApp extends StatelessWidget {
  const MainApp({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (context) => AuthBloc(),
        ),
        BlocProvider(
          create: (context) => BoxBloc(),
        ),
        BlocProvider(
          create: (context) => UserBloc(),
        ),
        BlocProvider(
          create: (context) => ControlBloc(),
        ),
        BlocProvider(
          create: (context) => DisplayBloc(),
        ),
        BlocProvider(
          create: (context) => DisplayDataBloc(),
        ),
      ],
      child: AuthGuard(child: HomeScreen()),
    );
  }
}
