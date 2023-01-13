import 'dart:convert';

import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/models/display_data_models.dart';
import 'package:http/http.dart' as http;

class DisplayDataBloc extends Bloc<MyEvent, MyState> {
  DisplayDataBloc() : super(StateInitialized()) {
    on<GetDisplayDataEvent>(_onGetDisplayData);
  }
  final _storage = const FlutterSecureStorage();

  Future _onGetDisplayData(GetDisplayDataEvent event, Emitter emit) async {
    try {
      emit(StateLoading());
      final token = await _storage.read(key: 'accessToken');
      final res = await http.get(
        Uri.parse(
            'https://api.rocket-translate.com/boxes/${event.boxId}/displays/${event.displayId}/data'),
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
      final displayData = json.map((e) => DisplayDatum.fromJson(e)).toList();
      emit(GetDisplayDataStateSuccess(displayData));
    } catch (e) {
      emit(StateError(message: 'something was wrong.'));
    }
  }
}

class GetDisplayDataEvent extends MyEvent {
  final int displayId;
  final String boxId;

  GetDisplayDataEvent({required this.displayId, required this.boxId});
}

class GetDisplayDataStateSuccess extends MyState {
  final List<DisplayDatum> displayData;

  GetDisplayDataStateSuccess(this.displayData);
}
