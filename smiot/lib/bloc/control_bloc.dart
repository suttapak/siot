import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import 'package:smiot/bloc/common_bloc.dart';
import 'package:http/http.dart' as http;
import 'package:smiot/models/control_models.dart';

class ControlBloc extends Bloc<MyEvent, MyState> {
  ControlBloc() : super(StateInitialized()) {
    on<GetControlEvent>(_onGetControls);
  }

  final _storage = const FlutterSecureStorage();

  Future<void> _onGetControls(GetControlEvent event, Emitter emit) async {
    try {
      emit(StateLoading());
      final token = await _storage.read(key: 'accessToken');
      final res = await http.get(
        Uri.parse(
            'https://api.rocket-translate.com/boxes/${event.boxId}/controls'),
        headers: <String, String>{
          'context-type': 'application/json',
          'Authorization': 'Bearer $token'
        },
      );
      if (res.statusCode != 200) {
        emit(StateError(message: res.body));
        return;
      }
      final json = jsonDecode(res.body) as List;
      final controls = json.map((e) => Controls.fromJson(e)).toList();
      emit(GetControlStateSuccess(controls));
    } catch (e) {
      emit(StateError(message: 'get display error'));
    }
  }
}

class GetControlEvent extends MyEvent {
  final String boxId;

  GetControlEvent(this.boxId);
}

class GetControlStateSuccess extends MyState {
  final List<Controls> controls;

  GetControlStateSuccess(this.controls);
}
